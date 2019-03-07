// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"
	"strings"

	client "github.com/bketelsen/lxdev/lxd"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list containers",
	Aliases: []string{"ls"},
	Long:    `List containers and their status.`,
	Run: func(cmd *cobra.Command, args []string) {
		lxclient, err := client.NewClient(socket)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
