// Copyright © 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"

	client "devlx/lxd"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List templates",
	Long:  `List available templates`,
	Run: func(cmd *cobra.Command, args []string) {

		lxclient, err := client.NewClient(config.LxdSocket)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}
		names, err := lxclient.ImageList()
		if err != nil {
			log.Error("Error executing command: " + err.Error())
			os.Exit(1)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name"})

		for _, name := range names {
			for _, alias := range name.Aliases {

				table.Append([]string{alias.Name})
			}

		}
		if len(names) < 1 {

			table.Append([]string{"{None Found}"})
		}
		table.Render()

	},
}

func init() {
	templateCmd.AddCommand(lsCmd)
}
