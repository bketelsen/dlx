package dlx

import (
	"os"

	"github.com/bketelsen/dlx/state"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type CmdProjectList struct {
	Cmd    *cobra.Command
	Global *state.Global
}

func (c *CmdProjectList) Command() *cobra.Command {
	// plistCmd represents the plist command
	var plistCmd = &cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

		RunE: c.Run,
	}
	return plistCmd
}
func (c *CmdProjectList) Run(cmd *cobra.Command, args []string) error {

	var err error
	cfg, err = getDlxConfig()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	conf := c.Global.Conf

	// Connect to the daemon
	remote, _, err := conf.ParseRemote(args[0])
	if err != nil {
		return err
	}

	d, err := conf.GetInstanceServer(remote)
	if err != nil {
		return err
	}
	projects, err := d.GetProjects()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defaultProject := c.Global.Conf.Remotes[c.Global.Conf.DefaultRemote].Project
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Description"})

	for _, project := range projects {
		if project.Name == defaultProject {
			project.Name = "*" + project.Name
		} else {
			project.Name = " " + project.Name
		}
		table.Append([]string{project.Name, project.Description})

	}
	if len(projects) < 1 {

		table.Append([]string{"{None Found}", "", ""})
	}
	table.Render()
	return nil

}
