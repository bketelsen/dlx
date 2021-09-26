// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package lxd

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	"github.com/bketelsen/dlx/config"
	"github.com/bketelsen/libgo/events"
	"github.com/buger/goterm"
	client "github.com/lxc/lxd/client"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/lxc/lxd/shared/termios"
	"github.com/pkg/errors"
)

type Container struct {
	Name      string
	Etag      string
	conn      client.ContainerServer
	container *api.Container
}
type sourceFile struct {
	path        string
	mode        int
	destination string
	filetype    string //"file or directory"
}

// GetContainer returns the container with the given `name`.
func GetContainer(conn client.ContainerServer, name string) (*Container, error) {
	container, etag, err := conn.GetContainer(name)
	if err != nil {
		return &Container{}, errors.Wrap(err, "getting container")
	}
	return &Container{
		container: container,
		conn:      conn,
		Name:      name,
		Etag:      etag,
	}, nil
}

// Stop causes a running container to cease running.
// An error is returned if the container is not running,
// or if the container doesn't exist.
func (c *Container) Stop() error {
	events.Publish(NewContainerState(c.Name, Stopping))
	cs := api.ContainerStatePut{
		Action: "stop",
	}
	op, err := c.conn.UpdateContainerState(c.Name, cs, c.Etag)
	if err != nil {
		return errors.Wrap(err, "updating container state")
	}
	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "waiting for container stop")
	}
	events.Publish(NewContainerState(c.Name, Stopped))
	return nil
}

// Start causes a stopped container to begin running.
// An error is returned if the container doesn't exist,
// or if the container is already running.
func (c *Container) Start() error {
	events.Publish(NewContainerState(c.Name, Starting))
	cs := api.ContainerStatePut{
		Action: "start",
	}
	op, err := c.conn.UpdateContainerState(c.Name, cs, c.Etag)
	if err != nil {
		return errors.Wrap(err, "starting container")
	}
	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "waiting for container start")
	}
	events.Publish(NewContainerState(c.Name, Started))
	return nil
}

// Remove deletes a stopped container.  An error is returned
// if the container is not stopped, or if the container doesn't
// exist.
func (c *Container) Remove() error {
	events.Publish(NewContainerState(c.Name, Removing))
	op, err := c.conn.DeleteContainer(c.Name)
	if err != nil {
		return errors.Wrap(err, "deleting container")
	}
	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "waiting for container delete")
	}

	events.Publish(NewContainerState(c.Name, Removed))
	return nil
}

func (c *Container) Exec(user string, command string, interactive bool) error {
	events.Publish(NewExecState(c.Name, command, Starting))
	terminalHeight := goterm.Height()
	terminalWidth := goterm.Width()
	// Setup the exec request
	environ := make(map[string]string)
	environ["TERM"] = os.Getenv("TERM")
	req := api.ContainerExecPost{
		Command:     []string{"/bin/bash", "-c", "sudo --user " + user + " --login" + " " + command},
		WaitForWS:   true,
		Interactive: interactive,
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
	op, err := c.conn.ExecContainer(c.Name, req, &largs)
	if err != nil {
		errors.Wrap(err, "execution error")
	}

	events.Publish(NewExecState(c.Name, command, Started))
	// Wait for it to complete
	err = op.Wait()
	if err != nil {
		errors.Wrap(err, "error waiting for execution")
	}

	events.Publish(NewExecState(c.Name, command, Completed))
	return nil
}

func (c *Container) Provision(user string) error {
	events.Publish(NewContainerState(c.Name, Provisioning))

	// this will fail if cloud-init isn't done.
	// need to make sure it's completed before running copykeys
	// TODO
	err := c.CopyKeys(user)
	if err != nil {
		return errors.Wrap(err, "copying ssh keys")
	}

	events.Publish(NewContainerState(c.Name, Provisioned))
	return nil
}

func (c *Container) Snapshot(snapshotName string) error {

	post := api.ContainerSnapshotsPost{
		Name: snapshotName,
	}
	op, err := c.conn.CreateContainerSnapshot(c.Name, post)

	if err != nil {
		return errors.Wrap(err, "create container snapshot")
	}
	// Wait for it to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "wait for operation")
	}
	return nil
}

func (c *Container) CopyFile(file sourceFile) error {
	events.Publish(NewCopyState(c.Name, file.destination, Started))
	var f *os.File
	var err error

	args := lxd.ContainerFileArgs{}
	if file.filetype == "file" {

		f, err = os.Open(file.path)
		if err != nil {
			return errors.New("Opening source file:" + err.Error())
		}

		defer f.Close()
		bb, err := ioutil.ReadAll(f)
		if err != nil {
			return errors.New("Reading source file:" + err.Error())
		}
		args = lxd.ContainerFileArgs{
			UID:       1000,
			GID:       1000,
			Content:   bytes.NewReader(bb),
			Type:      file.filetype,
			Mode:      file.mode,
			WriteMode: "overwrite",
		}
	} else {
		args = lxd.ContainerFileArgs{
			UID: 1000,
			GID: 1000,
			//	Content:   bytes.NewReader(bb),
			Type:      file.filetype,
			Mode:      file.mode,
			WriteMode: "overwrite",
		}
	}
	err = c.conn.CreateContainerFile(c.Name, file.destination, args)

	if err != nil {
		return errors.New("Creating destination file:" + err.Error())
	}

	events.Publish(NewCopyState(c.Name, file.destination, Completed))
	return nil
}

func (c *Container) CopyKeys(user string) error {
	files := []sourceFile{
		{path: filepath.Join(config.GetHomePath(), ".ssh"), mode: 0700, destination: "/home/" + user + "/.ssh", filetype: "directory"},
		{path: filepath.Join(config.GetHomePath(), ".ssh", "id_rsa.pub"), mode: 0644, destination: "/home/" + user + "/.ssh/id_rsa.pub", filetype: "file"},
		{path: filepath.Join(config.GetHomePath(), ".ssh", "id_rsa.pub"), mode: 0644, destination: "/home/" + user + "/.ssh/authorized_keys", filetype: "file"},
		{path: filepath.Join(config.GetHomePath(), ".ssh", "id_rsa"), mode: 0600, destination: "/home/" + user + "/.ssh/id_rsa", filetype: "file"},
	}

	for _, file := range files {
		err := c.CopyFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}
