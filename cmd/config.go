// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"

	"github.com/bketelsen/dlx/config"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage global configurations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		log.Running("Configuration")
		createConfig, err := cmd.Flags().GetBool("create")
		if err != nil {
			log.Error("Error getting flags: " + err.Error())
			os.Exit(1)
		}

		if createConfig {
			err := config.Create(cmd)
			if err != nil {
				log.Error("Error creating config file: " + err.Error())
				os.Exit(1)
			}
			log.Success("Default configuration file created")
			os.Exit(0)
		}
		cfg, err = config.Get()

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	configCmd.Flags().BoolP("create", "c", false, "Create global config file")
	configCmd.Flags().StringP("remote", "r", "", "LXD host network name or IP")
	configCmd.Flags().StringP("user", "u", "ubuntu", "Container username")
	configCmd.Flags().StringP("baseimage", "b", "dlxbase", "Default base image for new containers")
	configCmd.Flags().StringP("clientcert", "t", "", "Path to client certificate")
	configCmd.Flags().StringP("clientkey", "k", "", "Path to client key")
	configCmd.Flags().StringArrayP("profiles", "p", []string{}, "Profiles to use")
	configCmd.Flags().StringP("sshkey", "s", "", "Path to ssh private key authorized for HOST")
}
