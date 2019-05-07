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

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a template",
	Long:  `Remove a previously configured template.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		template := args[0]

		//Connect to LXD over the Unix Socket
		lxclient, err := client.NewClient(config.lxdSocket)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}

		// Remove the template which is an LXC image
		log.Running("Try to remove template: " + template)
		err = removeTemplate(lxclient, template)
		if err != nil {
			log.Error("Unable to remove template: " + err.Error())
			os.Exit(1)
		}
		log.Success("Template removed")
	},
}

func init() {
	templateCmd.AddCommand(rmCmd)
}
