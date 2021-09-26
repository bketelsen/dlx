// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop a running container",
	Long:  `Stop a running container.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		cfg, lxclient, err = connect()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		name = args[0]

		log.Running("Stopping container " + name)
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
