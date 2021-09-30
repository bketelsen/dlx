package dlx

import (
	"errors"
	"strings"

	"github.com/bketelsen/dlx/state"
	"github.com/spf13/cobra"
)

type CmdProjectSwitch struct {
	Cmd    *cobra.Command
	Global *state.Global
}

func (c *CmdProjectSwitch) Command() *cobra.Command {
	// pswitchCmd represents the pswitch command
	var pswitchCmd = &cobra.Command{
		Use:   "switch [project]",
		Short: "Switch to a different lxc project",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Args: cobra.MinimumNArgs(1),
		RunE: c.Run,
	}
	return pswitchCmd
}
func (c *CmdProjectSwitch) Run(cmd *cobra.Command, args []string) error {

	var err error
	cfg, err = getDlxConfig()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	conf := c.Global.Conf

	d, err := conf.GetInstanceServer(c.Global.Conf.DefaultRemote)
	if err != nil {
		return err
	}

	log.Running("Switching default project to: " + args[0])

	validProjects, err := d.GetProjects()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	var valid bool
	for _, project := range validProjects {
		if project.Name == args[0] {
			valid = true
		}
	}
	if !valid {
		log.Error("Project not found")
		var projects []string
		for _, project := range validProjects {
			projects = append(projects, project.Name)
		}

		log.Info("Valid Projects: " + strings.Join(projects, ", "))
		return errors.New("Project not found")
	}
	rem := conf.Remotes[conf.DefaultRemote]
	rem.Project = args[0]
	err = conf.SaveConfig(conf.ConfigPath("config.yml"))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Success("Switched default project to: " + args[0])
	return nil

}
