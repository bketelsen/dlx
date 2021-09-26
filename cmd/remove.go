// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "remove a container",
	Aliases: []string{"rm", "delete"},
	Long:    `Remove deletes a container.  It will fail if the container is running.`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		cfg, lxclient, err = connect()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		name = args[0]

		log.Running("Removing container " + name)

		err = lxclient.ContainerRemove(name)
		if err != nil {
			log.Error("Error executing command: " + err.Error())
			os.Exit(1)
		}

		log.Success("Removed container " + name)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
