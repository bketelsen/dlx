package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

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
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		cfg, lxclient, err = connect()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		log.Running("Switching default project to: " + args[0])

		validProjects, err := lxclient.GetProjects()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
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
			os.Exit(1)
		}
		err = lxcconf.SetDefaultProject(args[0])
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		log.Success("Switched default project to: " + args[0])

	},
}

func init() {
	projectCmd.AddCommand(pswitchCmd)
}
