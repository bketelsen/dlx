// Copyright © 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize your computer for devlx usage",
	Long: `The init command creates a configuration file, builds templates for
your containers, and creates lxc profiles that are required for operation.`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Running("Initializing devlx")

		log.Running("Creating configuration")
		err := initConfigFile()
		if err != nil {
			log.Error("Error initilizing config file: " + err.Error())
			os.Exit(1)
		}

		err = viper.WriteConfig()
		if err != nil {
			log.Error("Error initilizing config file: " + err.Error())
			os.Exit(1)
		}

		log.Success("Default configuration created")

		// create provision directory

		log.Running("Creating default provisioners")
		err = initDefaultProvisioners()
		if err != nil {
			log.Error("Error creating templates: " + err.Error())
			os.Exit(1)
		}
		log.Success("Provisioners created")

		// create templates directory

		log.Running("Creating profile templates")
		err = initProfileTemplates()
		if err != nil {
			log.Error("Error creating templates: " + err.Error())
			os.Exit(1)
		}
		log.Success("Templates created")

		// create lxc profiles

		log.Running("Creating lxc profiles")
		err = createProfiles("")
		if err != nil {
			log.Error("Error creating profiles: " + err.Error())
			os.Exit(1)
		}
		log.Success("Profiles created")
		// create relation store

		log.Running("Creating relation directory")
		err = createRelationsStore()
		if err != nil {
			log.Error("Error creating templates store: " + err.Error())
			os.Exit(1)
		}
		log.Success("Relation directory created")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
