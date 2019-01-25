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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

var (
	w bool
	s bool
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile [name]",
	Short: "create or replace the provisioning profile for lxdev",
	Long: `Profile creates or replaces the 'gui', 'cli', and 'util' profiles in lxc that allows you
to connect to running containers and possibly display X11 applications on the host.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name = args[0]
		log.Running("Managing profile " + name)
		c, err := lxd.ConnectLXDUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}

		exists := true
		prof, etag, err := c.GetProfile(name)
		if err != nil {
			exists = false
		}
		if w {
			filename := name + ".yaml"
			home, err := homedir.Dir()
			if err != nil {
				log.Error("Create Profile : " + err.Error())
				os.Exit(1)
			}
			fpath := filepath.Join(home, ".lxdev", "profiles", filename)
			f, err := os.Open(fpath)
			defer f.Close()
			if err != nil {
				log.Error("Create Profile : " + err.Error())
				log.Error("Try running `lxdev config -t` to create the templates directory.")
				os.Exit(1)
			}
			bb, err := ioutil.ReadAll(f)
			if err != nil {
				log.Error("Reading Profile : " + err.Error())
				os.Exit(1)
			}
			if exists {

				log.Running("Updating profile " + name)
				var profile api.ProfilePut
				err = yaml.Unmarshal(bb, &profile)
				if err != nil {
					log.Error("Parsing Profile : " + err.Error())
					os.Exit(1)
				}
				err = c.UpdateProfile(name, profile, etag)
				if err != nil {
					log.Error("Create Profile : " + err.Error())
					os.Exit(1)
				}

				log.Success("Updating profile " + name)
			} else {

				log.Running("Creating profile " + name)
				var profile api.ProfilesPost
				err = yaml.Unmarshal(bb, &profile)
				if err != nil {
					log.Error("Parsing Profile : " + err.Error())
					os.Exit(1)
				}
				profile.Name = name
				err = c.CreateProfile(profile)
				if err != nil {
					log.Error("Create Profile : " + err.Error())
					os.Exit(1)
				}
				log.Success("Creating profile " + name)
			}
		}

		if s {
			fmt.Println(prof, name)
		}

		log.Success("Managing profile " + name)
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	profileCmd.PersistentFlags().String("ethernet", "", "the name of your ethernet device e.g. 'enp5s0'")
	viper.BindPFlag("ethernet", profileCmd.PersistentFlags().Lookup("ethernet"))

	profileCmd.PersistentFlags().BoolVarP(&w, "write", "w", false, "Create or update a profile")
	viper.BindPFlag("write", profileCmd.PersistentFlags().Lookup("write"))

	profileCmd.PersistentFlags().BoolVarP(&s, "show", "l", false, "Show a profile")
	viper.BindPFlag("show", profileCmd.PersistentFlags().Lookup("show"))
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
