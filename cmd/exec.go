package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:     "exec [container] '[commands here]'",
	Aliases: []string{"run"},
	Short:   "Execute a command in a container",
	Args:    cobra.MinimumNArgs(2),
	Long: `Executes a command in the named container.  The command should be enclosed in 
single quotes.  e.g. exec mycontainer 'ls -la'`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		cfg, lxclient, err = connect()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		name = args[0]
		// Connect to LXD over the Unix socket
		err = lxclient.ContainerExec(name, strings.Join(args[1:], " "))
		if err != nil {
			log.Error("Error executing command: " + err.Error())
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
