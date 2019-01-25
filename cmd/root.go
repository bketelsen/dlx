// Copyright © 2019 Brian Ketelsen mail@bjk.fyi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/bketelsen/libgo/events"
	"github.com/dixonwille/wlog"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var log wlog.UI
var verbose bool
var socket string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lxdev",
	Short: "Provision lxd containers for development",
	Long:  `lxdev provisions lxd containers for local development.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lxdev.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging")

	rootCmd.PersistentFlags().StringVarP(&socket, "socket", "s", "/var/snap/lxd/common/lxd/unix.socket", "LXD Daemon socket")

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
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".lxdev" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".lxdev")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			log.Info("Using config file: " + viper.ConfigFileUsed())
		}
	}
}
