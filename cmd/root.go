// Copyright © 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"devlx/path"

	"github.com/bketelsen/libgo/events"
	"github.com/dixonwille/wlog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var log wlog.UI
var verbose bool

var config devlxConfig

var rootCmd = &cobra.Command{
	Use:   "devlx",
	Short: "Provision lxd containers for development",
	Long:  `devlx provisions lxd containers for local development.`,
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
	rootCmd.PersistentFlags().StringVarP(&config.LxdSocket, "lxd-socket", "s", viper.GetString("lxd-socket"), "LXD Daemon socket")

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

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name "devlx" (without extension).
		viper.AddConfigPath(path.GetConfigPath())
		viper.SetConfigName("devlx")
	}

	viper.AutomaticEnv()

	viper.SetDefault("ssh-socket", "/run/user/1234/keyring/ssh")
	viper.SetDefault("display", ":0")

	err := viper.BindEnv("ssh-socket", "SSH_AUTH_SOCK")
	if err != nil {
		log.Error(err.Error())
	}

	err = viper.BindEnv("DISPLAY")
	if err != nil {
		log.Error(err.Error())
	}

	// If a config file is found, read it in.
	if err = viper.ReadInConfig(); err == nil {
		if verbose {
			log.Info("Using config file: " + viper.ConfigFileUsed())
		}
	} else if err != nil {
		fmt.Println("No configuration files found.")

		log.Error("Run `devlx init` to setup devlx config.")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Error("Error unmarshling config")
		os.Exit(1)
	}

}
