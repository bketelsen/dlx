// Copyright © 2019 Brian Ketelsen mail@bjk.fyi
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
	"github.com/spf13/viper"

	client "github.com/bketelsen/lxdev/lxd"
)

var (
	name     string
	template string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a container",
	Long:  `Create a new container from a template.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name = args[0]
		log.Running("Creating container " + name)
		// Connect to LXD over the Unix socket
		// TODO: account for non snap install
		lxclient, err := client.NewClient(socket)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}
		err = lxclient.ContainerCreate(name, true, template, getProfiles())
		if err != nil {
			log.Error("Unable to create container: " + err.Error())
			os.Exit(1)
		}
		log.Success("Created container " + name)
	},
}

func getProfiles() []string {
	profiles := []string{}

	// default gui if nothing set
	// probably a way to do this with some sort of set in viper/pflag?  TODO
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	createCmd.PersistentFlags().StringVar(&template, "template", "", "base template for container")
	viper.BindPFlag("template", createCmd.PersistentFlags().Lookup("template"))

}
