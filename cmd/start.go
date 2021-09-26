// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a paused container",
	Long:  `Start a paused container.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		cfg, lxclient, err = connect()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		name = args[0]

		log.Running("Starting container " + name)
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
