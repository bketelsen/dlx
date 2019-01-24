package lxd

import (
	"os"
	"syscall"

	"github.com/bketelsen/libgo/events"
	"github.com/buger/goterm"
	client "github.com/lxc/lxd/client"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/lxc/lxd/shared/termios"
	"github.com/pkg/errors"
)

type Connection struct {
	URL string

	conn client.ContainerServer
}

// NewConnection creates a new connection to an LXD Daemon,
// returning a Connection
func NewConnection(url string) (*Connection, error) {
	c := &Connection{
		URL: url,
	}
	err := c.Connect()
	return c, err
}

// Connect establishes a connection to an LXD Daemon
func (c *Connection) Connect() error {
	var err error
	c.conn, err = client.ConnectLXDUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
	if err != nil {
		return errors.Wrap(err, "Error connecting to LXD daemon")
	}

	events.Publish(NewConnectionCreated(c))
	return nil
}

func (c *Connection) Shell(name string) error {

	terminalHeight := goterm.Height()
	terminalWidth := goterm.Width()
	// Setup the exec request
	environ := make(map[string]string)
	environ["TERM"] = os.Getenv("TERM")
	// TODO: Make the command for this configurable?
	req := api.ContainerExecPost{
		Command:     []string{"/bin/bash", "-c", "sudo --user ubuntu --login"},
		WaitForWS:   true,
		Interactive: true,
		Width:       terminalWidth,
		Height:      terminalHeight,
		Environment: environ,
	}

	// Setup the exec arguments (fds)
	largs := lxd.ContainerExecArgs{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	// Setup the terminal (set to raw mode)
	if req.Interactive {
		cfd := int(syscall.Stdin)
		oldttystate, err := termios.MakeRaw(cfd)
		if err != nil {
			return errors.Wrap(err, "unable to make raw terminal")
		}

		defer termios.Restore(cfd, oldttystate)
	}

	// Get the current state
	op, err := c.conn.ExecContainer(name, req, &largs)
	if err != nil {
		return errors.Wrap(err, "error calling exec")
	}

	// Wait for it to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "error waiting for completion")
		return err
	}
	return nil
}

func (c *Connection) Create(name string) error {

	// Container creation request
	req := api.ContainersPost{
		Name: name,
		ContainerPut: api.ContainerPut{
			//Profiles: getProfiles(),
			Profiles: []string{"default"},
		},
		Source: api.ContainerSource{
			Type: "image",
			//Server:   "https://cloud-images.ubuntu.com/daily",
			//Alias:    getImage(),
			Alias: "clibase",
			//Protocol: "simplestreams",
		},
	}

	events.Publish(NewContainerState(name, Creating))

	// Get LXD to create the container (background operation)
	op, err := c.conn.CreateContainer(req)
	if err != nil {
		return errors.Wrap(err, "creating container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "wait for create container")
	}

	events.Publish(NewContainerState(name, Created))
	// Get LXD to start the container (background operation)
	reqState := api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}

	events.Publish(NewContainerState(name, Starting))
	op, err = c.conn.UpdateContainerState(name, reqState, "")
	if err != nil {
		return errors.Wrap(err, "starting container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "waiting for container start")
	}

	events.Publish(NewContainerState(name, Started))
	return nil
}
func (c *Connection) Exec(name string, command string) error {

	events.Publish(NewExecState(name, command, Starting))
	terminalHeight := goterm.Height()
	terminalWidth := goterm.Width()
	// Setup the exec request
	environ := make(map[string]string)
	environ["TERM"] = os.Getenv("TERM")
	req := api.ContainerExecPost{
		Command:     []string{"/bin/bash", "-c", "sudo --user ubuntu --login" + " " + command},
		WaitForWS:   true,
		Interactive: false,
		Width:       terminalWidth,
		Height:      terminalHeight,
		Environment: environ,
	}

	// Setup the exec arguments (fds)
	largs := lxd.ContainerExecArgs{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	// Setup the terminal (set to raw mode)
	if req.Interactive {
		cfd := int(syscall.Stdin)
		oldttystate, err := termios.MakeRaw(cfd)
		if err != nil {
			return errors.Wrap(err, "error making raw terminal")
		}

		defer termios.Restore(cfd, oldttystate)
	}

	// Get the current state
	op, err := c.conn.ExecContainer(name, req, &largs)
	if err != nil {
		errors.Wrap(err, "execution error")
	}

	events.Publish(NewExecState(name, command, Started))
	// Wait for it to complete
	err = op.Wait()
	if err != nil {
		errors.Wrap(err, "error waiting for execution")
	}

	events.Publish(NewExecState(name, command, Completed))
	return nil
}
