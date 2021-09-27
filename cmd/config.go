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
			if verbose {
				log.Info("Creating dlx config file")
			}
			err := config.Create(cmd)
			if err != nil {
				log.Error("Error creating config file: " + err.Error())
				os.Exit(1)
			}
			log.Success("Default configuration file created")
			os.Exit(0)
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolP("create", "c", false, "Create global config file")
	configCmd.Flags().StringP("user", "u", "ubuntu", "Container username")
	configCmd.Flags().StringP("baseimage", "b", "dlxbase", "Default base image for new containers")
	configCmd.Flags().StringArrayP("profiles", "p", []string{}, "Profiles to use")
	configCmd.Flags().StringP("sshkey", "s", "", "Path to ssh private key authorized for HOST")
}
