// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package dlx

import (
	"net"

	"io/ioutil"

	"github.com/bketelsen/dlx/state"
	"github.com/lxc/lxd/shared/api"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"golang.org/x/crypto/ssh"
)

var (
	name string
)

type CmdCreate struct {
	Global *state.Global
	name   string
}

// createCmd represents the create command
func (c *CmdCreate) Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create [name]",
		Short: "Create a container",
		Long:  `Create a new container.`,
		Args:  cobra.MinimumNArgs(1),
		RunE:  c.Run,
	}
	cmd.Flags().StringP("baseimage", "b", "", "(optional) base image to use")

	return cmd
}

func (c *CmdCreate) Run(cmd *cobra.Command, args []string) error {
	conf := c.Global.Conf

	// Quick checks.
	exit, err := c.Global.CheckArgs(cmd, args, 1, 1)
	if exit {
		return err
	}

	// Connect to LXD
	remote, name, err := conf.ParseRemote(args[0])
	if err != nil {
		return err
	}

	d, err := conf.GetInstanceServer(remote)
	if err != nil {
		return err
	}

	name = args[0]

	var bi string
	baseimg, err := cmd.Flags().GetString("baseimage")
	if err != nil {
		log.Error("Error getting flags: " + err.Error())
		return err
	}
	if baseimg != "" {
		bi = baseimg
	} else {
		bi = cfg.Remotes[c.Global.Conf.DefaultRemote].BaseImage
	}
	//err = lxclient.ContainerCreate(name, true, bi, []string{"default"})
	var source api.ContainerSource

	source = api.ContainerSource{
		Type: "image",
		//Server:   "https://cloud-images.ubuntu.com/daily",
		//Alias:    getImage(),
		Alias: bi,
		//Protocol: "simplestreams",
	}

	req := api.ContainersPost{
		Name: name,
		ContainerPut: api.ContainerPut{
			Profiles: []string{"default"},
		},
		Source: source,
	}

	// Get LXD to create the container (background operation)
	op, err := d.CreateContainer(req)
	if err != nil {
		return errors.Wrap(err, "creating container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "wait for create container")
	}

	// Get LXD to start the container (background operation)
	reqState := api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}

	op, err = d.UpdateContainerState(name, reqState, "")
	if err != nil {
		return errors.Wrap(err, "starting container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "waiting for container start")
	}

	log.Success("Created container " + name)

	log.Running("Provisioning container " + name)

	key, err := ioutil.ReadFile(cfg.Remotes[c.Global.Conf.DefaultRemote].SSHPrivateKey)
	if err != nil {
		log.Error("unable to read private key" + err.Error())
		return err
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Error("unable to parse private key" + err.Error())
	}

	config := &ssh.ClientConfig{
		User: cfg.Remotes[c.Global.Conf.DefaultRemote].User,
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	host := c.Global.Conf.DefaultRemote

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
	return err
}
