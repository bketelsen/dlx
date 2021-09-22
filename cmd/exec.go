// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"
	"strings"

	client "github.com/bketelsen/dlx/lxd"
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
		err := getConfig()
		if err != nil {
			log.Error("Unable to get configuration:" + err.Error())
		}
		name = args[0]
		// Connect to LXD over the Unix socket
		lxclient, err := client.NewClient(cfg)
		if err != nil {
			log.Error("Unable to connect: " + err.Error())
			os.Exit(1)
		}
		err = lxclient.ContainerExec(name, strings.Join(args[1:], " "))
		if err != nil {
			log.Error("Error executing command: " + err.Error())
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
