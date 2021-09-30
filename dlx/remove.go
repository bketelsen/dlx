// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package dlx

import (
	"github.com/bketelsen/dlx/state"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type CmdRemove struct {
	Cmd    *cobra.Command
	Global *state.Global
}

func (c *CmdRemove) Command() *cobra.Command {

	// removeCmd represents the remove command
	var removeCmd = &cobra.Command{
		Use:     "remove",
		Short:   "remove a container",
		Aliases: []string{"rm", "delete"},
		Long:    `Remove deletes a container.  It will fail if the container is running.`,
		Args:    cobra.MinimumNArgs(1),
		RunE:    c.Run,
	}
	return removeCmd
}
func (c *CmdRemove) Run(cmd *cobra.Command, args []string) error {
	conf := c.Global.Conf

	// Quick checks.
	exit, err := c.Global.CheckArgs(cmd, args, 1, -1)
	if exit {
		return err
	}

	d, err := conf.GetInstanceServer(c.Global.Conf.DefaultRemote)
	if err != nil {
		return err
	}

	name = args[0]

	log.Running("Removing container " + name)

	op, err := d.DeleteInstance(name)
	if err != nil {
		return errors.Wrap(err, "deleting container")
	}
	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "waiting for container delete")
	}

	return nil
	log.Success("Removed container " + name)
	return nil
}
