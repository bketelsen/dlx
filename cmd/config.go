// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"devlx/path"
	"net"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"errors"
	"os"
	"path/filepath"
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

func getConfigValues() error {

	//Sets default UID if $UID is not se in env. (This should be on most linux systems)
	viper.SetDefault("uid", 1000)

	//Find and set LXD Unix socket.
	lxdSocket := ""
	possibleLxdSockets := []string{
		"/var/snap/lxd/common/lxd/unix.socket",
		"/var/lib/lxd/unix.socket",
	}

	for _, socket := range possibleLxdSockets {
		if _, err := os.Stat(socket); err == nil {
			lxdSocket = socket
			break
		}
	}

	viper.Set("lxdSocket", lxdSocket)

	//Get Network adapters
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	var interfaceNames []string

	for _, inter := range interfaces {

		if strings.Contains(inter.Name, "lo") || strings.Contains(inter.Name, "tun") || strings.Contains(inter.Name, "docker") {
			continue
		}

		interfaceNames = append(interfaceNames, inter.Name)
	}

	if l := len(interfaceNames); l > 1 {
		prompt := promptui.Select{
			Label: "Select Network Adapter",
			Items: interfaceNames,
		}

		_, result, err := prompt.Run()

		if err != nil {
			return err
		}

		viper.Set("ethernet", result)
	} else if l == 1 {
		viper.Set("ethernet", interfaceNames[0])
	} else {
		//in the future we should probably create a host network like docker does.
		return errors.New("No network interfaces available")
	}

	return nil

}

func createConfig() error {
	//make config directory and file
	err := os.MkdirAll(filepath.Join(path.GetConfigPath()), 0755)

	// f, err := os.Create(filepath.Join(path.GetConfigPath(), "devlx.yaml"))
	err = viper.WriteConfig()
	if err != nil {
		return err
	}
	// _, err = f.Write([]byte(configTemplate))
	// if err != nil {
	// 	return err
	// }
	return nil
}

func checkConfig() error {
	_, err := os.Stat(filepath.Join(path.GetConfigPath(), "devlx.yaml"))
	if err != nil {
		return err
	}
	return nil
}

func createTemplates() error {
	box := packr.New("provision", "../templates/provision")

	err := os.MkdirAll(filepath.Join(path.GetConfigPath(), "provision"), 0755)
	if err != nil {
		return err
	}
	for _, tpl := range box.List() {
		bb, err := box.Find(tpl)
		if err != nil {
			return err
		}
		f, err := os.Create(filepath.Join(path.GetConfigPath(), "provision", tpl))
		if err != nil {
			return err
		}
		_, err = f.Write([]byte(bb))
		if err != nil {
			return err
		}
	}

	pbox := packr.New("profiles", "../templates/profiles")
	err = os.MkdirAll(filepath.Join(path.GetConfigPath(), "profiles"), 0755)
	if err != nil {
		return err
	}
	for _, tpl := range pbox.List() {
		bb, err := pbox.Find(tpl)
		if err != nil {
			return err
		}
		f, err := os.Create(filepath.Join(path.GetConfigPath(), "profiles", tpl))
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

	err := os.MkdirAll(filepath.Join(path.GetConfigPath(), "templates"), 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(path.GetConfigPath(), "templates", "relations.yaml"))
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

// const configTemplate = `
// cliimage: "18.10"
// guiimage: "18.10"
// utilimage: "18.10"
// ethernet: enp5s0
// `
