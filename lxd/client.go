package lxd

import (
	"github.com/bketelsen/libgo/events"
	client "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/pkg/errors"
)

type Client struct {
	URL string

	conn client.ContainerServer
}

// NewClient creates a new connection to an LXD Daemon,
// returning a Client
func NewClient(url string) (*Client, error) {
	c := &Client{
		URL: url,
	}
	err := c.Connect()
	return c, err
}

// Connect establishes a connection to an LXD Daemon
func (c *Client) Connect() error {
	var err error
	c.conn, err = client.ConnectLXDUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
	if err != nil {
		return errors.Wrap(err, "Error connecting to LXD daemon")
	}

	events.Publish(NewConnectionCreated(c))
	return nil
}

func (c *Client) ContainerShell(name string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Exec("", true)
}

func (c *Client) ContainerCreate(name string) error {
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
func (c *Client) ContainerExec(name string, command string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Exec(command, false)
}

func (c *Client) ContainerList() ([]string, error) {
	names, err := c.conn.GetContainerNames()
	if err != nil {
		errors.Wrap(err, "get container names")
	}
	return names, err
}

func (c *Client) ContainerInfo(name string) (*api.Container, error) {
	container, _, err := c.conn.GetContainer(name)
	if err != nil {
		errors.Wrap(err, "get container names")
	}
	return container, err
}

func (c *Client) ContainerRemove(name string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Remove()
}

func (c *Client) ContainerStart(name string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Start()
}

func (c *Client) ContainerStop(name string) error {
	cont, err := GetContainer(c.conn, name)
	if err != nil {
		return errors.Wrap(err, "getting container")
	}
	return cont.Stop()
}
