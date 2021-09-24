// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"bytes"
	"net"
	"os"

	"io/ioutil"

	"github.com/spf13/cobra"

	"golang.org/x/crypto/ssh"

	client "github.com/bketelsen/dlx/lxd"
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

		err := getConfig()
		if err != nil {
			log.Error("Unable to get configuration:" + err.Error())
		}
		name = args[0]
		log.Running("Creating container " + name)
		lxclient, err := client.NewClient(cfg)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}

        var bi string
		baseimg, err := cmd.Flags().GetString("baseimage")
		if err != nil {
			log.Error("Error getting flags: " + err.Error())
			os.Exit(1)
		}
		if baseimg != "" {
		    bi = baseimg
		} else {
		bi = cfg.BaseImage}
		err = lxclient.ContainerCreate(name, true, bi, cfg.Profiles)
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

		key, err := ioutil.ReadFile(cfg.SSHPrivateKey)
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
			User: cfg.User,
			Auth: []ssh.AuthMethod{
				// Use the PublicKeys method for remote authentication.
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}

		// Connect to the remote server and perform the SSH handshake.
		client, err := ssh.Dial("tcp", cfg.Host+":22", config)
		if err != nil {
			log.Error("unable to connect:" + err.Error())
		}
		defer client.Close()
		newSession, err := client.NewSession()
		if err != nil {
			log.Error("unable to connect:" + err.Error())
		}

		var b bytes.Buffer
		newSession.Stdout = &b
		defer newSession.Close()
		//lxc config device add $container dlxbind disk source=$HOME/projects/$container path=/home/`whoami`/projects/$container
		if err := newSession.Run("devices " + name); err != nil {
			log.Error("unable to run command:" + err.Error())
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
