package dlx

import (
	"os"

	"github.com/bketelsen/dlx/state"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type CmdRemoteList struct {
	Cmd    *cobra.Command
	Global *state.Global
}

func (c *CmdRemoteList) Command() *cobra.Command {
	// rlistCmd represents the rlist command
	var rlistCmd = &cobra.Command{
		Use:   "list",
		Short: "List remote LXC servers",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: c.Run,
	}
	return rlistCmd
}
func (c *CmdRemoteList) Run(cmd *cobra.Command, args []string) error {

	conf := c.Global.Conf

	remotes := conf.Remotes
	defremote := conf.DefaultRemote

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Address"})

	for name, remote := range remotes {
		if name == defremote {
			name = "*" + name
		} else {
			name = " " + name
		}
		table.Append([]string{name, remote.Addr})

	}
	if len(remotes) < 1 {

		table.Append([]string{"{None Found}", "", ""})
	}
	table.Render()
	return nil
}
