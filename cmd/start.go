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

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a paused container",
	Long:  `Start a paused container.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name = args[0]

		log.Running("Starting container " + name)
		lxclient, err := client.NewClient(config.lxdSocket)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}
		err = lxclient.ContainerStart(name)
		if err != nil {
			log.Error("Error executing command: " + err.Error())
			os.Exit(1)
		}
		log.Success("Container " + name + " started.")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
