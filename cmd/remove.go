// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"

	client "github.com/bketelsen/dlx/lxd"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "remove a container",
	Aliases: []string{"rm"},
	Long:    `Remove deletes a container.  It will fail if the container is running.`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		err := getConfig()
		if err != nil {
			log.Error("Unable to get configuration:" + err.Error())
		}

		name = args[0]

		log.Running("Removing container " + name)

		lxclient, err := client.NewClient(cfg)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}
		err = lxclient.ContainerRemove(name)
		if err != nil {
			log.Error("Error executing command: " + err.Error())
			os.Exit(1)
		}

		log.Success("Removed container " + name)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
