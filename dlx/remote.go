package dlx

import (
	"github.com/bketelsen/dlx/state"
	"github.com/spf13/cobra"
)

type CmdRemote struct {
	Global *state.Global
}

func (c *CmdRemote) Command() *cobra.Command {
	// remoteCmd represents the remote command
	var remoteCmd = &cobra.Command{
		Use:   "remote",
		Short: "View and manage LXC remotes",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}

	rlistCmd := CmdRemoteList{Cmd: remoteCmd, Global: c.Global}
	remoteCmd.AddCommand(rlistCmd.Command())

	rswitchCmd := CmdRemoteList{Cmd: remoteCmd, Global: c.Global}
	remoteCmd.AddCommand(rswitchCmd.Command())
	return remoteCmd
}
