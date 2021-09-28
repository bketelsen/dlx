package dlx

import (
	"strings"

	"github.com/bketelsen/dlx/state"
	"github.com/spf13/cobra"
)

type CmdRemoteSwitch struct {
	Cmd    *cobra.Command
	Global *state.Global
}

func (c *CmdRemoteSwitch) Command() *cobra.Command {
	// switchCmd represents the switch command
	var switchCmd = &cobra.Command{
		Use:   "switch",
		Short: "Change LXC remote server",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: c.Run,
	}
	return switchCmd
}
func (c *CmdRemoteSwitch) Run(cmd *cobra.Command, args []string) error {
	var err error

	log.Running("Switching remote server to: " + args[0])

	conf := c.Global.Conf

	validRemotes := conf.Remotes

	var valid bool
	for name, _ := range validRemotes {
		if name == args[0] {
			valid = true
		}
	}
	if !valid {
		log.Error("Remote not found")
		var remotes []string
		for name, _ := range validRemotes {
			remotes = append(remotes, name)
		}

		log.Info("Valid Remotes: " + strings.Join(remotes, ", "))
		return err
	}
	conf.DefaultRemote = args[0]
	err = conf.SaveConfig(conf.ConfigPath("config.yml"))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Success("Switched default remote to: " + args[0])
	return nil
}
