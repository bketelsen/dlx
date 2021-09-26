package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// plistCmd represents the plist command
var plistCmd = &cobra.Command{
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
		projects, err := lxclient.GetProjects()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		defaultProject := lxcconf.DefaultRemote().Project
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Description"})

		for _, project := range projects {
			if project.Name == defaultProject {
				project.Name = "*" + project.Name
			}
			table.Append([]string{project.Name, project.Description})

		}
		if len(projects) < 1 {

			table.Append([]string{"{None Found}", "", ""})
		}
		table.Render()

	},
}

func init() {
	projectCmd.AddCommand(plistCmd)
}
