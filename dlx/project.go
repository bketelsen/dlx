package dlx

import (
	"github.com/bketelsen/dlx/state"
	"github.com/spf13/cobra"
)

type CmdProject struct {
	Global *state.Global
	Cmd    *cobra.Command
}

func (c *CmdProject) Command() *cobra.Command {
	// projectCmd represents the project command
	var projectCmd = &cobra.Command{
		Use:   "project",
		Short: "View and manage lxc projects",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}

	plistCmd := CmdProjectList{Cmd: projectCmd, Global: c.Global}
	projectCmd.AddCommand(plistCmd.Command())

	pswitchCmd := CmdProjectSwitch{Cmd: projectCmd, Global: c.Global}
	projectCmd.AddCommand(pswitchCmd.Command())
	return projectCmd
}
