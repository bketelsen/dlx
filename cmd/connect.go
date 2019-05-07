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

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:     "connect",
	Aliases: []string{"shell"},
	Short:   "connect to a running container",
	Long:    `Connect to a running container.`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name = args[0]

		log.Running("Connecting to container " + name)
		// Connect to LXD over the Unix Socket
		lxclient, err := client.NewClient(config.lxdSocket)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}

		err = lxclient.ContainerShell(name)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}

		log.Success("Closed connection to container " + name)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
