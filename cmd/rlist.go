package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// rlistCmd represents the rlist command
var rlistCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		cfg, lxclient, err = connect()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		remotes := lxcconf.GetRemotes()
		defremote := lxcconf.Config.DefaultRemote

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Address"})

		for name, remote := range remotes {
			if name == defremote {
				name = "*" + name
			}
			table.Append([]string{name, remote.Addr})

		}
		if len(remotes) < 1 {

			table.Append([]string{"{None Found}", "", ""})
		}
		table.Render()
	},
}

func init() {
	remoteCmd.AddCommand(rlistCmd)
}
