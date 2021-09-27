package cmd

import (
	"fmt"
	"os"

	"github.com/bketelsen/dlx/config"
	client "github.com/bketelsen/dlx/lxd"
	"github.com/dixonwille/wlog"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cfg *config.Config
var lxclient *client.Client
var lxcconf *config.LXC
var log wlog.UI
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dlx",
	Short: "Provision lxd containers for development",
	Long: `dlx provisions lxd containers for local development.
See https://dlx.rocks for full documentation.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	log = wlog.New(os.Stdin, os.Stdout, os.Stderr)

	log = wlog.AddPrefix("?", wlog.Cross, "i", "-", ">", "~", wlog.Check, "!", log)
	log = wlog.AddConcurrent(log)
	log = wlog.AddColor(wlog.None, wlog.Red, wlog.Blue, wlog.None, wlog.None, wlog.None, wlog.Cyan, wlog.Green, wlog.Magenta, log)

}

func connect() (*config.Config, *client.Client, error) {
	if verbose {
		log.Info(fmt.Sprintf("Verbose : %v", verbose))
	}
	var err error
	cfg, err = config.Get()
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to get configuration")
	}

	lxcconf, err = LXCConfig()
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to get LXC configuration")
	}

	// Connect to LXD over the Unix socket
	lxclient, err := client.NewClient(cfg, lxcconf)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to connect")
	}
	defaultProject := lxcconf.DefaultRemote().Project
	log.Info("Connected to " + lxcconf.Config.DefaultRemote + "/" + defaultProject)

	return cfg, lxclient, err

}

func LXCConfig() (*config.LXC, error) {

	lxcconf, err := config.GetLXCConfig()
	if err != nil {
		log.Error(err.Error())
		log.Info("Install the lxd command line app and connect to your lxd server before running dlx.")
		return nil, err
	}
	return lxcconf, err
}
