// Copyright © 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"

	client "devlx/lxd"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "remove a container",
	Aliases: []string{"rm"},
	Long:    `Remove deletes a container.  It will fail if the container is running.`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name = args[0]

		log.Running("Removing container " + name)

		lxclient, err := client.NewClient(config.LxdSocket)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}
		err = lxclient.ContainerRemove(name)
		if err != nil {
			log.Error("Error executing command: " + err.Error())
			os.Exit(1)
		}

		err = setContainerTemplateRelation(lxclient, name, "", false)
		if err != nil {
			log.Error("Error destroy image relation : " + err.Error())
			os.Exit(1)
		}

		log.Success("Removed container " + name)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
