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
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	name        string
	skipkeys    bool
	skipdefault bool
	clionly     bool
	gui         bool
	util        bool
)
var guiimage string
var cliimage string
var utilimage string

var guiinit string
var cliinit string
var utilinit string

var extraprofiles *[]string

type sourceFile struct {
	path        string
	mode        int
	destination string
	filetype    string //"file or directory"
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a container",
	Long: `Create a new container for a project, provisioned with your preferred 
development tools.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name = args[0]
		// Connect to LXD over the Unix socket
		// TODO: account for non snap install
		c, err := lxd.ConnectLXDUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
		if err != nil {
			log.Fatal("Connect:", err)
		}

		// Container creation request
		req := api.ContainersPost{
			Name: name,
			ContainerPut: api.ContainerPut{
				Profiles: getProfiles(),
			},
			Source: api.ContainerSource{
				Type:     "image",
				Server:   "https://cloud-images.ubuntu.com/daily",
				Alias:    getImage(),
				Protocol: "simplestreams",
			},
		}
		log.Printf("Creating container with image: %s, profile(s): %v\n", getImage(), getProfiles())
		// Get LXD to create the container (background operation)
		op, err := c.CreateContainer(req)
		if err != nil {
			log.Fatal("Create:", err)
		}

		// Wait for the operation to complete
		err = op.Wait()
		if err != nil {
			log.Fatal("Wait:", err)
		}

		// Get LXD to start the container (background operation)
		reqState := api.ContainerStatePut{
			Action:  "start",
			Timeout: -1,
		}

		op, err = c.UpdateContainerState(name, reqState, "")
		if err != nil {
			log.Fatal("UpdateState:", err)
		}

		// Wait for the operation to complete
		err = op.Wait()
		if err != nil {
			log.Fatal("Wait:", err)
		}

		if !skipkeys {
			err = copyFiles(c, name)
			if err != nil {
				log.Fatal("Copy Files:", err)
			}
		}
	},
}

func getProfiles() []string {
	profiles := []string{}

	// default gui if nothing set
	// probably a way to do this with some sort of set in viper/pflag?  TODO
	c := viper.GetBool("clionly")
	g := viper.GetBool("gui")
	u := viper.GetBool("util")

	if !c && !g && !u {
		g = true
	}
	if !viper.GetBool("skipdefault") {
		profiles = append(profiles, "default")
	}
	if c {
		profiles = append(profiles, "cli")
	}
	if u {
		profiles = append(profiles, "util")
	}
	if g {
		profiles = append(profiles, "gui")
	}

	for _, prof := range *extraprofiles {
		profiles = append(profiles, prof)
	}
	return profiles
}
func getImage() string {
	if clionly {
		return viper.GetString("cliimage")
	}
	if util {
		return viper.GetString("utilimage")
	}
	return viper.GetString("guiimage")
}
func copyFiles(c lxd.ContainerServer, name string) error {
	// HACK: Find out when provisioning is done??
	time.Sleep(5 * time.Second)

	files := []sourceFile{
		sourceFile{path: "/home/bketelsen/.ssh", mode: 0700, destination: "/home/ubuntu/.ssh", filetype: "directory"},
		sourceFile{path: "/home/bketelsen/.ssh/id_rsa.pub", mode: 0644, destination: "/home/ubuntu/.ssh/id_rsa.pub", filetype: "file"},
		sourceFile{path: "/home/bketelsen/.ssh/id_rsa", mode: 0600, destination: "/home/ubuntu/.ssh/id_rsa", filetype: "file"},
	}

	for _, file := range files {
		var f *os.File
		var err error
		log.Printf("[Creating] %s\n", file.destination)

		args := lxd.ContainerFileArgs{}
		if file.filetype == "file" {

			f, err = os.Open(file.path)
			defer f.Close()
			if err != nil {
				return errors.New("Opening source file:" + err.Error())
			}
			bb, err := ioutil.ReadAll(f)
			if err != nil {
				return errors.New("Reading source file:" + err.Error())
			}
			args = lxd.ContainerFileArgs{
				UID:       1000,
				GID:       1000,
				Content:   bytes.NewReader(bb),
				Type:      file.filetype,
				Mode:      file.mode,
				WriteMode: "overwrite",
			}
		} else {
			args = lxd.ContainerFileArgs{
				UID: 1000,
				GID: 1000,
				//	Content:   bytes.NewReader(bb),
				Type:      file.filetype,
				Mode:      file.mode,
				WriteMode: "overwrite",
			}
		}
		err = c.CreateContainerFile(name, file.destination, args)

		if err != nil {
			return errors.New("Creating destination file:" + err.Error())
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().BoolVar(&skipkeys, "skipkeys", false, "Skip copy of ssh keys")
	viper.BindPFlag("skipkeys", createCmd.Flags().Lookup("skipkeys"))

	createCmd.Flags().BoolVar(&skipdefault, "skipdefault", false, "Skip default profile")
	viper.BindPFlag("skipdefault", createCmd.Flags().Lookup("skipdefault"))
	createCmd.Flags().BoolVar(&clionly, "cli", false, "Use CLI-only profile, no X Windows support.")
	viper.BindPFlag("clionly", createCmd.Flags().Lookup("cli"))

	createCmd.Flags().BoolVar(&gui, "gui", false, "Use GUI profile, with X Windows support. (default)")
	viper.BindPFlag("gui", createCmd.Flags().Lookup("gui"))

	createCmd.Flags().BoolVar(&util, "util", false, "Use UTIL profile, with X Windows support.")
	viper.BindPFlag("util", createCmd.Flags().Lookup("util"))

	createCmd.PersistentFlags().StringVar(&guiimage, "guiimage", "18.10", "Ubuntu version for GUI instances")
	viper.BindPFlag("guiimage", createCmd.PersistentFlags().Lookup("guiimage"))

	createCmd.PersistentFlags().StringVar(&cliimage, "cliimage", "18.10", "Ubuntu version for CLI instances")
	viper.BindPFlag("cliimage", createCmd.PersistentFlags().Lookup("cliimage"))

	createCmd.PersistentFlags().StringVar(&utilimage, "utilimage", "18.10", "Ubuntu version for UTIL instances")
	viper.BindPFlag("utilimage", createCmd.PersistentFlags().Lookup("utilimage"))

	createCmd.PersistentFlags().StringVar(&guiinit, "guiinit", "go.yaml", "cloud-init for GUI instances")
	viper.BindPFlag("guiinit", createCmd.PersistentFlags().Lookup("guiinit"))

	createCmd.PersistentFlags().StringVar(&cliinit, "cliinit", "go.yaml", "cloud-init for CLI instances")
	viper.BindPFlag("cliinit", createCmd.PersistentFlags().Lookup("cliinit"))

	createCmd.PersistentFlags().StringVar(&utilinit, "utilinit", "go.yaml", "cloud-init for UTIL instances")
	viper.BindPFlag("utilinit", createCmd.PersistentFlags().Lookup("utilinit"))

	extraprofiles = createCmd.PersistentFlags().StringSlice("profiles", []string{}, "Comma separated list of extra profiles to add. e.g. 'go,neovim'")

}
