// Copyright © 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"
	"strings"

	client "devlx/lxd"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list containers",
	Aliases: []string{"ls"},
	Long:    `List containers and their status.`,
	Run: func(cmd *cobra.Command, args []string) {
		lxclient, err := client.NewClient(config.LxdSocket)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}
		names, err := lxclient.ContainerList()
		if err != nil {
			log.Error("Error executing command: " + err.Error())
			os.Exit(1)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Status", "Profile(s)"})

		for _, name := range names {
			container, err := lxclient.ContainerInfo(name)
			if err != nil {
				log.Error("Get Container: " + err.Error())
				os.Exit(1)
			}
			table.Append([]string{container.Name, container.Status, strings.Join(container.Profiles, ",")})

		}
		if len(names) < 1 {

			table.Append([]string{"{None Found}", "", ""})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
