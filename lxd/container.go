package lxd

import (
	"github.com/bketelsen/libgo/events"
	client "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/pkg/errors"
)

type Container struct {
	Name      string
	Etag      string
	conn      client.ContainerServer
	container *api.Container
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
