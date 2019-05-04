// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"devlx/config"
	"devlx/path"

	"github.com/bketelsen/libgo/events"
	"github.com/dixonwille/wlog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var log wlog.UI
var verbose bool
var socket string
var devlxConfig config.GlobalConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devlx",
	Short: "Provision lxd containers for development",
	Long:  `devlx provisions lxd containers for local development.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		err := checkConfig()
		if err != nil {
			fmt.Println("No configuration files found.")

			log.Error("Run `devlx init` to create required lxc profiles.")
		} else {
			cmd.Help()
		}
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", filepath.Join(path.GetConfigPath(), "devlx.yaml"), "path to config file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging")
	rootCmd.PersistentFlags().StringVarP(&socket, "socket", "s", viper.GetString("socket"), "LXD Daemon socket")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	log = wlog.New(os.Stdin, os.Stdout, os.Stderr)

	log = wlog.AddPrefix("?", wlog.Cross, "i", "-", ">", "~", wlog.Check, "!", log)
	log = wlog.AddConcurrent(log)
	log = wlog.AddColor(wlog.None, wlog.Red, wlog.Blue, wlog.None, wlog.None, wlog.None, wlog.Cyan, wlog.Green, wlog.Magenta, log)

	eh := &events.Subscriber{
		Handler: eventHandler,
	}
	// Subscribe to an Event
	events.Subscribe(eh)

}

func eventHandler(e events.Event) {
	switch t := e.(type) {
	//case *TopicStart:
	// check the topic and create if needed
	//	fmt.Println(t.Topic)
	default:
		// we don't care
		if verbose {
			log.Info(fmt.Sprintf("\t%T\t %s\n", t, e.Name()))
		}
	}

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name "devlx" (without extension).
		viper.AddConfigPath(path.GetConfigPath())
		viper.SetConfigName("devlx")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			log.Info("Using config file: " + viper.ConfigFileUsed())
		}
	}

	err := viper.Unmarshal(&devlxConfig)
	if err != nil {
		panic("Unable to unmarshal config")
	}
}
