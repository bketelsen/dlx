// Copyright © 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	client "devlx/lxd"
)

var (
	name     string
	template string
)

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a container",
	Long:  `Create a new container from a template.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name = args[0]
		log.Running("Creating container " + name)
		lxclient, err := client.NewClient(config.lxdSocket)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}
		err = lxclient.ContainerCreate(name, true, template, getProfiles())
		if err != nil {
			log.Error("Unable to create container: " + err.Error())
			os.Exit(1)
		}

		// Store the LXC Image -> Container relationship
		log.Running("Storing image container relation ")
		err = setContainerTemplateRelation(lxclient, name, template, true)
		if err != nil {
			log.Error("Unable to create container-template relations" + err.Error())
			os.Exit(1)
		}

		log.Success("Created container " + name)
	},
}

func getProfiles() []string {
	profiles := []string{}

	// default gui if nothing set
	c := viper.GetBool("clionly")
	g := viper.GetBool("gui")

	if !c && !g {
		g = true
	}
	if !viper.GetBool("skipdefault") {
		profiles = append(profiles, "default")
	}
	if c {
		profiles = append(profiles, "cli")
	}
	if g {
		profiles = append(profiles, "gui")
	}

	return profiles
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringVar(&template, "template", viper.GetString("template"), "base template for container")
	// viper.BindPFlag("template", createCmd.PersistentFlags().Lookup("template"))

}
