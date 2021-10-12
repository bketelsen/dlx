// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package dlx

import (
	"github.com/bketelsen/dlx/state"
	"github.com/lxc/lxd/shared/api"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

var (
	name string
)

type CmdCreate struct {
	Global *state.Global
	name   string
}

// createCmd represents the create command
func (c *CmdCreate) Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create [name]",
		Short: "Create a container",
		Long:  `Create a new container.`,
		Args:  cobra.MinimumNArgs(1),
		RunE:  c.Run,
	}
	cmd.Flags().StringP("baseimage", "b", "", "(optional) base image to use")

	return cmd
}

func (c *CmdCreate) Run(cmd *cobra.Command, args []string) error {
	conf := c.Global.Conf

	// Quick checks.
	exit, err := c.Global.CheckArgs(cmd, args, 1, 1)
	if exit {
		return err
	}

	// Connect to LXD
	remote, _, err := conf.ParseRemote(args[0])
	if err != nil {
		return err
	}

	d, err := conf.GetInstanceServer(remote)
	if err != nil {
		return err
	}

	name = args[0]

	var bi string
	baseimg, err := cmd.Flags().GetString("baseimage")
	if err != nil {
		log.Error("Error getting flags: " + err.Error())
		return err
	}
	if baseimg != "" {
		bi = baseimg
	} else {
		bi = cfg.Remotes[c.Global.Conf.DefaultRemote].BaseImage
	}
	//err = lxclient.ContainerCreate(name, true, bi, []string{"default"})

	source := api.ContainerSource{
		Type: "image",
		//Server:   "https://cloud-images.ubuntu.com/daily",
		//Alias:    getImage(),
		Alias: bi,
		//Protocol: "simplestreams",
	}

	req := api.ContainersPost{
		Name: name,
		ContainerPut: api.ContainerPut{
			Profiles: []string{"default"}, // TODO: ? support command line adding profiles
		},
		Source: source,
	}

	// Get LXD to create the container (background operation)
	op, err := d.CreateContainer(req)
	if err != nil {
		return errors.Wrap(err, "creating container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "wait for create container")
	}

	// Get LXD to start the container (background operation)
	reqState := api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}

	op, err = d.UpdateContainerState(name, reqState, "")
	if err != nil {
		return errors.Wrap(err, "starting container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "waiting for container start")
	}

	log.Success("Created container " + name)

	return err
}
