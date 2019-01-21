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
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/spf13/cobra"
)

var (
	lxcpath    string
	template   string
	distro     string
	release    string
	arch       string
	name       string
	verbose    bool
	flush      bool
	validation bool
)

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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
				Profiles: []string{"default", "gui"},
			},
			Source: api.ContainerSource{
				Type: "image",
				//Server:   "https://cloud-images.ubuntu.com/daily",
				Alias: "brian",
				//Protocol: "simplestreams",
			},
		}

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
		err = copyFiles(c, name)
		if err != nil {
			log.Fatal("Copy Files:", err)
		}
	},
}

func copyFiles(c lxd.ContainerServer, name string) error {
	// HACK: Find out when provisioning is done??
	time.Sleep(30 * time.Second)

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
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
