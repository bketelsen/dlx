package state

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"

	"github.com/bketelsen/dlx/config"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared"
	cli "github.com/lxc/lxd/shared/cmd"
	"github.com/lxc/lxd/shared/i18n"
	"github.com/lxc/lxd/shared/logger"
	"github.com/lxc/lxd/shared/logging"
	"github.com/lxc/lxd/shared/version"
	"github.com/spf13/cobra"
)

type Global struct {
	Conf     *config.Config
	ConfPath string
	Cmd      *cobra.Command
	Ret      int

	FlagForceLocal bool
	FlagHelp       bool
	FlagHelpAll    bool
	FlagLogDebug   bool
	FlagLogVerbose bool
	FlagProject    string
	FlagQuiet      bool
	FlagVersion    bool
}

func (c *Global) PreRun(cmd *cobra.Command, args []string) error {
	var err error
	// If calling the help, skip pre-run
	if cmd.Name() == "help" {
		return nil
	}

	// Figure out the config directory and config path
	var configDir string
	if os.Getenv("LXD_CONF") != "" {
		configDir = os.Getenv("LXD_CONF")
	} else if os.Getenv("HOME") != "" {
		configDir = path.Join(os.Getenv("HOME"), ".config", "lxc")
	} else {
		user, err := user.Current()
		if err != nil {
			return err
		}

		configDir = path.Join(user.HomeDir, ".config", "lxc")
	}

	c.ConfPath = os.ExpandEnv(path.Join(configDir, "config.yml"))

	// Load the configuration
	if c.FlagForceLocal {
		c.Conf = config.NewConfig("", true)
	} else if shared.PathExists(c.ConfPath) {
		c.Conf, err = config.LoadConfig(c.ConfPath)
		if err != nil {
			return err
		}
	} else {
		c.Conf = config.NewConfig(filepath.Dir(c.ConfPath), true)
	}

	// Override the project
	if c.FlagProject != "" {
		c.Conf.ProjectOverride = c.FlagProject
	}

	// Setup password helper
	c.Conf.PromptPassword = func(filename string) (string, error) {
		return cli.AskPasswordOnce(fmt.Sprintf(i18n.G("Password for %s: "), filename)), nil
	}

	// If the user is running a command that may attempt to connect to the local daemon
	// and this is the first time the client has been run by the user, then check to see
	// if LXD has been properly configured.  Don't display the message if the var path
	// does not exist (LXD not installed), as the user may be targeting a remote daemon.
	if !c.FlagForceLocal && shared.PathExists(shared.VarPath("")) && !shared.PathExists(c.ConfPath) {
		// Create the config dir so that we don't get in here again for this user.
		err = os.MkdirAll(c.Conf.ConfigDir, 0750)
		if err != nil {
			return err
		}

		// And save the initial configuration
		err = c.Conf.SaveConfig(c.ConfPath)
		if err != nil {
			return err
		}

		// Attempt to connect to the local server
		runInit := true
		d, err := lxd.ConnectLXDUnix("", nil)
		if err == nil {
			info, _, err := d.GetServer()
			if err == nil && info.Environment.Storage != "" {
				runInit = false
			}
		}

		flush := false
		if runInit {
			fmt.Fprintf(os.Stderr, i18n.G("If this is your first time running LXD on this machine, you should also run: lxd init")+"\n")
			flush = true
		}

		if !shared.StringInSlice(cmd.Name(), []string{"init", "launch"}) {
			fmt.Fprintf(os.Stderr, i18n.G("To start your first instance, try: lxc launch ubuntu:18.04")+"\n")
			flush = true
		}

		if flush {
			fmt.Fprintf(os.Stderr, "\n")
		}
	}

	// Set the user agent
	c.Conf.UserAgent = version.UserAgent

	// Setup the logger
	logger.Log, err = logging.GetLogger("", "", c.FlagLogVerbose, c.FlagLogDebug, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Global) PostRun(cmd *cobra.Command, args []string) error {
	// Macaroon teardown
	if c.Conf != nil && shared.PathExists(c.ConfPath) {
		// Save cookies on exit
		c.Conf.SaveCookies()
	}

	return nil
}
func (c *Global) ParseServers(remotes ...string) ([]RemoteResource, error) {
	servers := map[string]lxd.InstanceServer{}
	resources := []RemoteResource{}

	for _, remote := range remotes {
		// Parse the remote
		remoteName, name, err := c.Conf.ParseRemote(remote)
		if err != nil {
			return nil, err
		}

		// Setup the struct
		resource := RemoteResource{
			Remote: remoteName,
			Name:   name,
		}

		// Look at our cache
		_, ok := servers[remoteName]
		if ok {
			resource.Server = servers[remoteName]
			resources = append(resources, resource)
			continue
		}

		// New connection
		d, err := c.Conf.GetInstanceServer(remoteName)
		if err != nil {
			return nil, err
		}

		resource.Server = d
		servers[remoteName] = d
		resources = append(resources, resource)
	}

	return resources, nil
}

func (c *Global) CheckArgs(cmd *cobra.Command, args []string, minArgs int, maxArgs int) (bool, error) {
	if len(args) < minArgs || (maxArgs != -1 && len(args) > maxArgs) {
		cmd.Help()

		if len(args) == 0 {
			return true, nil
		}

		return true, fmt.Errorf(i18n.G("Invalid number of arguments"))
	}

	return false, nil
}

type RemoteResource struct {
	Remote string
	Server lxd.InstanceServer
	Name   string
}
