// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

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
				os.Exit(1)
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

func checkConfig() error {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	_, err = os.Stat(filepath.Join(home, ".lxdev.yaml"))
	if err != nil {
		return err
	}
	return nil
}
func createTemplates() error {
	box := packr.New("provision", "../templates/provision")
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(home, ".lxdev", "provision"), 0755)
	if err != nil {
		return err
	}
	for _, tpl := range box.List() {
		bb, err := box.Find(tpl)
		if err != nil {
			return err
		}
		f, err := os.Create(filepath.Join(home, ".lxdev", "provision", tpl))
		if err != nil {
			return err
		}
		_, err = f.Write([]byte(bb))
		if err != nil {
			return err
		}
	}

	pbox := packr.New("profiles", "../templates/profiles")
	err = os.MkdirAll(filepath.Join(home, ".lxdev", "profiles"), 0755)
	if err != nil {
		return err
	}
	for _, tpl := range pbox.List() {
		bb, err := pbox.Find(tpl)
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

func createRelationsStore() error {
	// create config storage directory and file
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(home, ".lxdev", "templates"), 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(home, ".lxdev", "templates", "relations.yaml"))
	if err != nil {
		return err
	}

	defer f.Close()

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
