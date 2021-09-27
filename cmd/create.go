// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"net"
	"os"

	"io/ioutil"

	"github.com/spf13/cobra"

	"golang.org/x/crypto/ssh"
)

var (
	name string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a container",
	Long:  `Create a new container.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		cfg, lxclient, err = connect()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		name = args[0]

		var bi string
		baseimg, err := cmd.Flags().GetString("baseimage")
		if err != nil {
			log.Error("Error getting flags: " + err.Error())
			os.Exit(1)
		}
		if baseimg != "" {
			bi = baseimg
		} else {
			bi = cfg.Remotes[lxcconf.Config.DefaultRemote].BaseImage
		}
		err = lxclient.ContainerCreate(name, true, bi, []string{"default"})
		if err != nil {
			log.Error("Unable to create container: " + err.Error())
			os.Exit(1)
		}

		log.Success("Created container " + name)

		log.Running("Provisioning container " + name)
		err = lxclient.ContainerProvision(name)

		if err != nil {
			log.Error("Unable to provision container: " + err.Error())
			os.Exit(1)
		}

		key, err := ioutil.ReadFile(cfg.Remotes[lxcconf.Config.DefaultRemote].SSHPrivateKey)
		if err != nil {
			log.Error("unable to read private key" + err.Error())
			os.Exit(1)
		}

		// Create the Signer for this private key.
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			log.Error("unable to parse private key" + err.Error())
		}

		config := &ssh.ClientConfig{
			User: cfg.Remotes[lxcconf.Config.DefaultRemote].User,
			Auth: []ssh.AuthMethod{
				// Use the PublicKeys method for remote authentication.
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}

		host := lxcconf.Config.DefaultRemote
		if verbose {
			log.Info("Connecting to " + host)
		}
		// Connect to the remote server and perform the SSH handshake.
		client, err := ssh.Dial("tcp", host+":22", config)
		if err != nil {
			log.Error("unable to connect:" + err.Error())
		}
		defer client.Close()
		newSession, err := client.NewSession()
		if err != nil {
			log.Error("unable to connect:" + err.Error())
		}

		defer newSession.Close()

		if verbose {
			log.Info("Running provisioning script")
		}
		//lxc config device add $container dlxbind disk source=$HOME/projects/$container path=/home/`whoami`/projects/$container
		output, err := newSession.CombinedOutput("/usr/local/bin/devices " + name)
		if err != nil {
			log.Error("unable to run command:" + err.Error())
			log.Error(string(output))
		}
		log.Success("Provisioned container " + name)
		newSession.Close()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	createCmd.Flags().StringP("baseimage", "b", "", "(optional) base image to use")

}
