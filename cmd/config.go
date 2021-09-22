// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bketelsen/dlx/config"
	"github.com/bketelsen/dlx/path"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage global configurations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		log.Running("Configuration")
		config, err := cmd.Flags().GetBool("create")
		if err != nil {
			log.Error("Error getting flags: " + err.Error())
			os.Exit(1)
		}

		if config {
			err := createConfig(cmd)
			if err != nil {
				log.Error("Error creating config file: " + err.Error())
				os.Exit(1)
			}
			log.Success("Default configuration file created")
			os.Exit(0)
		}
		err = checkConfig()
		if err != nil {
			log.Error("Error finding configuration file: " + err.Error())
			log.Warn("Run 'dlx config -c' to create a default configuration file")
			os.Exit(1)
		}
		log.Success("Configuration file found at" + filepath.Join(path.GetConfigPath(), "dlx.yaml"))
		bb, err := ioutil.ReadFile(filepath.Join(path.GetConfigPath(), "dlx.yaml"))
		if err != nil {
			log.Error("Error reading configuration file: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("--------")
		fmt.Println(string(bb))

	},
}

func createConfig(cmd *cobra.Command) error {
	//make config directory and file
	err := os.MkdirAll(filepath.Join(path.GetConfigPath()), 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(path.GetConfigPath(), "dlx.yaml"))
	if err != nil {
		return err
	}
	defer f.Close()

	host, err := cmd.Flags().GetString("remote")
	if err != nil {
		log.Error("Error getting host flag: " + err.Error())
	}

	user, err := cmd.Flags().GetString("user")
	if err != nil {
		log.Error("Error getting user flag: " + err.Error())
	}

	bi, err := cmd.Flags().GetString("baseimage")
	if err != nil {
		log.Error("Error getting baseimage flag: " + err.Error())
	}

	cc, err := cmd.Flags().GetString("clientcert")
	if err != nil {
		log.Error("Error getting clientcert flag: " + err.Error())
	}

	ck, err := cmd.Flags().GetString("clientkey")
	if err != nil {
		log.Error("Error getting clientkey flag: " + err.Error())
	}

	sshkey, err := cmd.Flags().GetString("sshkey")
	if err != nil {
		log.Error("Error getting sshkey flag: " + err.Error())
	}
	profs, err := cmd.Flags().GetStringArray("profiles")
	if err != nil {
		log.Error("Error getting clientkey flag: " + err.Error())
	}
	config := &config.Config{
		Host:          host,
		User:          user,
		Socket:        uri,
		BaseImage:     bi,
		ClientCert:    cc,
		ClientKey:     ck,
		Profiles:      profs,
		SSHPrivateKey: sshkey,
	}
	bb, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	_, err = f.Write(bb)
	if err != nil {
		return err
	}
	return nil
}

func checkConfig() error {
	_, err := os.Stat(filepath.Join(path.GetConfigPath(), "dlx.yaml"))
	if err != nil {
		return err
	}
	return nil
}

func getConfig() error {

	err := checkConfig()
	if err != nil {
		log.Warn("Run 'dlx config -c' to create a default configuration file")
		return err
	}

	bb, err := ioutil.ReadFile(filepath.Join(path.GetConfigPath(), "dlx.yaml"))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bb, &cfg)
	return err

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
	configCmd.Flags().StringP("baseimage", "b", "dlxbase", "Base image for new containers")
	configCmd.Flags().StringP("clientcert", "t", "", "Path to client certificate")
	configCmd.Flags().StringP("clientkey", "k", "", "Path to client key")
	configCmd.Flags().StringArrayP("profiles", "p", []string{}, "Profiles to use")
	configCmd.Flags().StringP("sshkey", "s", "", "Path to ssh private key authorized for HOST")
}
