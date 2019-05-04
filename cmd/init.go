// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize your computer for devlx usage",
	Long: `The init command creates a configuration file, builds templates for
your containers, and creates lxc profiles that are required for operation.`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Running("Initializing devlx")

		log.Running("Collecting Environment Info")
		err := getConfigValues()
		if err != nil {
			log.Error("Error getting config values: " + err.Error())
			os.Exit(1)
		}
		log.Success("Collection complete")

		log.Running("Creating configuration file")
		err = createConfig()
		if err != nil {
			log.Error("Error creating config file: " + err.Error())
			os.Exit(1)
		}
		log.Success("Default configuration file created")

		// create templates directory

		log.Running("Creating templates")
		err = createTemplates()
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
