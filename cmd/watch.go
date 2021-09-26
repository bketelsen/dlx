package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "watch",
	Long:  `watch`,
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		cfg, lxclient, err = connect()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		err = lxclient.Watch()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
