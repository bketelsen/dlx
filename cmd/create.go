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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/buger/goterm"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/lxc/lxd/shared/termios"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	name        string
	skipkeys    bool
	skipdefault bool
	clionly     bool
	gui         bool
)
var guiimage string
var cliimage string

var provisioners *[]string

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
		log.Running("Creating container " + name)
		// Connect to LXD over the Unix socket
		// TODO: account for non snap install
		c, err := lxd.ConnectLXDUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
		if err != nil {
			log.Error("Connect: " + err.Error())
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
		log.Info("Creating container with image: " + getImage() + " profile(s): " + strings.Join(getProfiles(), ","))
		// Get LXD to create the container (background operation)
		op, err := c.CreateContainer(req)
		if err != nil {
			log.Error("Create: " + err.Error())
			os.Exit(1)
		}

		// Wait for the operation to complete
		err = op.Wait()
		if err != nil {
			log.Error("Wait: " + err.Error())
			os.Exit(1)
		}

		// Get LXD to start the container (background operation)
		reqState := api.ContainerStatePut{
			Action:  "start",
			Timeout: -1,
		}

		op, err = c.UpdateContainerState(name, reqState, "")
		if err != nil {
			log.Error("Start: " + err.Error())
			os.Exit(1)
		}

		// Wait for the operation to complete
		err = op.Wait()
		if err != nil {
			log.Error("Wait: " + err.Error())
			os.Exit(1)
		}

		if !skipkeys {
			err = copyFiles(c, name)
			if err != nil {
				log.Error("Copy Files: " + err.Error())
				os.Exit(1)
			}
		}

		err = provision(c)
		if err != nil {
			log.Error("Provisioning: " + err.Error())
			os.Exit(1)
		}

		log.Success("Created container " + name)
	},
}

func provision(c lxd.ContainerServer) error {
	final := make([]string, 0)
	cli := viper.GetBool("clionly")
	gui := viper.GetBool("gui")

	if !cli && !gui {
		gui = true
	}
	if cli {
		final = append(final, "clibase")
	}

	if gui {
		final = append(final, "guibase")
	}

	final = append(final, *provisioners...)
	fmt.Println(final)
	fmt.Println(*provisioners)
	for _, prof := range final {
		home, err := homedir.Dir()
		if err != nil {
			log.Error("Provision Home Dir: " + err.Error())
			os.Exit(1)
		}
		file := sourceFile{
			path:        filepath.Join(home, ".lxdev", "provision", prof+".sh"),
			mode:        0755,
			destination: filepath.Join("/", "tmp", prof+".sh"),
			filetype:    "file",
		}

		// copy the file in
		err = copyFile(c, file)
		if err != nil {
			log.Error("Copy Provisioner: " + err.Error())
			os.Exit(1)
		}
		terminalHeight := goterm.Height()
		terminalWidth := goterm.Width()
		// Setup the exec request
		environ := make(map[string]string)
		environ["TERM"] = os.Getenv("TERM")
		req := api.ContainerExecPost{
			Command:     []string{"/bin/bash", "-c", "sudo --user ubuntu --login /bin/bash -c /tmp/" + prof + ".sh"},
			WaitForWS:   true,
			Interactive: true,
			Width:       terminalWidth,
			Height:      terminalHeight,
			Environment: environ,
		}

		// Setup the exec arguments (fds)
		largs := lxd.ContainerExecArgs{
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		// Setup the terminal (set to raw mode)
		if req.Interactive {
			cfd := int(syscall.Stdin)
			oldttystate, err := termios.MakeRaw(cfd)
			if err != nil {
				log.Error("Make Raw Terminal" + err.Error())

			}

			defer termios.Restore(cfd, oldttystate)
		}

		// Get the current state
		op, err := c.ExecContainer(name, req, &largs)
		if err != nil {
			log.Error("Exec: " + err.Error())
		}

		// Wait for it to complete
		err = op.Wait()
		if err != nil {
			log.Error("Wait: " + err.Error())
		}

	}
	return nil
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
func getImage() string {
	if clionly {
		return viper.GetString("cliimage")
	}
	return viper.GetString("guiimage")
}

func copyFile(c lxd.ContainerServer, file sourceFile) error {

	var f *os.File
	var err error
	log.Running("Creating " + file.destination)

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

	log.Success("Created " + file.destination)
	return nil
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
		err := copyFile(c, file)
		if err != nil {
			return err
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

	createCmd.PersistentFlags().StringVar(&guiimage, "guiimage", "18.10", "Ubuntu version for GUI instances")
	viper.BindPFlag("guiimage", createCmd.PersistentFlags().Lookup("guiimage"))

	createCmd.PersistentFlags().StringVar(&cliimage, "cliimage", "18.10", "Ubuntu version for CLI instances")
	viper.BindPFlag("cliimage", createCmd.PersistentFlags().Lookup("cliimage"))

	provisioners = createCmd.PersistentFlags().StringSlice("provisioners", []string{}, "Comma separated list of provision scripts to run . e.g. 'go,neovim'")

}
