// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"

	"github.com/spf13/cobra"

	client "github.com/bketelsen/dlx/lxd"
)

var (
	name     string
	template string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a container",
	Long:  `Create a new container.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		err := getConfig()
		if err != nil {
			log.Error("Unable to get configuration:" + err.Error())
		}
		name = args[0]
		log.Running("Creating container " + name)
		// Connect to LXD over the Unix socket
		// TODO: account for non snap install
		lxclient, err := client.NewClient(cfg)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}
		err = lxclient.ContainerCreate(name, true, cfg.BaseImage, cfg.Profiles)
		if err != nil {
			log.Error("Unable to create container: " + err.Error())
			os.Exit(1)
		}

		log.Success("Created container " + name)

		log.Running("Provisioning container " + name)
		err = lxclient.ContainerProvision(name)

		if err != nil {
			log.Error("Unable to provision container: " + err.Error())
			os.Exit(1)
		}

		log.Success("Provisioned container " + name)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}
