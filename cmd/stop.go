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

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop a running container",
	Long:  `Stop a running container.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name = args[0]

		log.Running("Stopping container " + name)
		lxclient, err := client.NewClient(config.LxdSocket)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}
		err = lxclient.ContainerStop(name)
		if err != nil {
			log.Error("Error stopping container: " + err.Error())
			os.Exit(1)
		}

		log.Success("Container " + name + " stopped.")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
