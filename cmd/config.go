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
	"path/filepath"

	"github.com/gobuffalo/packr/v2"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage global configurations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		log.Running("Create configuration")
		config, err := cmd.Flags().GetBool("create")
		if err != nil {
			log.Error("Error getting flags: " + err.Error())
			os.Exit(1)
		}

		templates, err := cmd.Flags().GetBool("templates")
		if err != nil {
			log.Error("Error getting flags: " + err.Error())
			os.Exit(1)
		}
		if !config && !templates {
			cmd.Usage()
			log.Error("Please specify either one or more flags.")
			os.Exit(1)
		}
		if config {
			err := createConfig()
			if err != nil {
				log.Error("Error creating config file: " + err.Error())
			}
			log.Success("Default configuration file created")
		}
		if templates {
			err := createTemplates()
			if err != nil {
				log.Error("Error creating templates: " + err.Error())
			}
			log.Success("Templates created")
		}

		log.Success("Configuration created")
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
	box := packr.New("templates", "../templates")
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(home, ".lxdev", "profiles"), 0755)
	if err != nil {
		return err
	}
	for _, tpl := range box.List() {
		bb, err := box.Find(tpl)
		if err != nil {
			return err
		}
		f, err := os.Create(filepath.Join(home, ".lxdev", "profiles", tpl))
		if err != nil {
			return err
		}
		_, err = f.Write([]byte(bb))
		if err != nil {
			return err
		}
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
guiimage: "18.10"
utilimage: "18.10"
ethernet: enp5s0`
