package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Change LXC remote server",
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

		log.Running("Switching remote server to: " + args[0])

		validRemotes := lxcconf.GetRemotes()

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
			os.Exit(1)
		}
		err = lxcconf.SetDefaultRemote(args[0])
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		log.Success("Switched default remote to: " + args[0])

	},
}

func init() {
	remoteCmd.AddCommand(switchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
