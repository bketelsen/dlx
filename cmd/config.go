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
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage global configurations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := cmd.Flags().GetBool("create")
		if err != nil {
			log.Fatal("Error getting flags:", err)
		}

		templates, err := cmd.Flags().GetBool("templates")
		if err != nil {
			log.Fatal("Error getting flags:", err)
		}
		if !config && !templates {
			cmd.Usage()
			log.Fatal("Please specify either one or more flags.")
		}
		if config {
			err := createConfig()
			if err != nil {
				log.Fatal("Error creating config file:", err)
			}
		}
		if templates {
			err := createTemplates()
			if err != nil {
				log.Fatal("Error creating templates:", err)
			}
		}
	},
}

func createConfig() error {

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(home, ".lxdev.yaml"))
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(configTemplate))
	if err != nil {
		return err
	}
	return nil
}

func createTemplates() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(home, ".lxdev", "profiles"), 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(home, ".lxdev", "cloud-init"), 0755)
	if err != nil {
		return err
	}
	return nil
}
func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	configCmd.Flags().BoolP("create", "c", false, "Create global config file in $HOME")
	configCmd.Flags().BoolP("templates", "t", false, "Create global template folders in $HOME")
}

const configTemplate = `cliimage: "18.10"
cliinit: go.yaml
guiimage: "18.10"
guiinit: go.yaml
utilimage: "18.10"
utilinit: go.yaml
ethernet: enp5s0`
