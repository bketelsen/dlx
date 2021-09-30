package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bketelsen/dlx/config"
	"github.com/bketelsen/dlx/state"

	"github.com/bketelsen/dlx/dlx"
	"github.com/bketelsen/dlx/lxc"
	"github.com/lxc/lxd/shared/i18n"
)

func main() {

	// "root" command, as a parent to the lxc command, plus the dlx commands
	app := &cobra.Command{}
	app.Use = "dlx"

	globalCmd := state.Global{Cmd: app}

	// Wrappers
	app.PersistentPreRunE = globalCmd.PreRun
	app.PersistentPostRunE = globalCmd.PostRun
	app.PersistentFlags().BoolVar(&globalCmd.FlagVersion, "version", false, i18n.G("Print version number"))
	app.PersistentFlags().BoolVarP(&globalCmd.FlagHelp, "help", "h", false, i18n.G("Print help"))
	app.PersistentFlags().BoolVar(&globalCmd.FlagForceLocal, "force-local", false, i18n.G("Force using the local unix socket"))
	app.PersistentFlags().StringVar(&globalCmd.FlagProject, "project", "", i18n.G("Override the source project")+"``")
	app.PersistentFlags().BoolVar(&globalCmd.FlagLogDebug, "debug", false, i18n.G("Show all debug messages"))
	app.PersistentFlags().BoolVarP(&globalCmd.FlagLogVerbose, "verbose", "v", false, i18n.G("Show all information messages"))
	app.PersistentFlags().BoolVarP(&globalCmd.FlagQuiet, "quiet", "q", false, i18n.G("Don't show progress information"))

	lxcCmd := lxc.CmdLxc{Global: &globalCmd}
	app.AddCommand(lxcCmd.Command())

	configCmd := dlx.CmdConfig{Global: &globalCmd}
	app.AddCommand(configCmd.Command())
	createCmd := dlx.CmdCreate{Global: &globalCmd}
	app.AddCommand(createCmd.Command())
	connectCmd := dlx.CmdConnect{Global: &globalCmd}
	app.AddCommand(connectCmd.Command())
	docsCmd := dlx.CmdDocs{Global: &globalCmd}
	app.AddCommand(docsCmd.Command())
	execCmd := dlx.CmdExec{Global: &globalCmd}
	app.AddCommand(execCmd.Command())
	listCmd := dlx.CmdList{Global: &globalCmd}
	app.AddCommand(listCmd.Command())
	projectCmd := dlx.CmdProject{Global: &globalCmd}
	app.AddCommand(projectCmd.Command())
	remoteCmd := dlx.CmdRemote{Global: &globalCmd}
	app.AddCommand(remoteCmd.Command())
	removeCmd := dlx.CmdRemove{Global: &globalCmd}
	app.AddCommand(removeCmd.Command())
	startCmd := dlx.CmdStart{Global: &globalCmd}
	app.AddCommand(startCmd.Command())
	stopCmd := dlx.CmdStop{Global: &globalCmd}
	app.AddCommand(stopCmd.Command())

	consoleCmd := dlx.CmdConsole{Global: &globalCmd}
	app.AddCommand(consoleCmd.Command())
	// Run the main command and handle errors
	err := app.Execute()
	if err != nil {
		// Handle non-Linux systems
		if err == config.ErrNotLinux {
			fmt.Fprintf(os.Stderr, i18n.G(`This client hasn't been configured to use a remote LXD server yet.
As your platform can't run native Linux instances, you must connect to a remote LXD server.

If you already added a remote server, make it the default with "lxc remote switch NAME".
To easily setup a local LXD server in a virtual machine, consider using: https://multipass.run`)+"\n")
			os.Exit(1)
		}

		// Default error handling
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if globalCmd.Ret != 0 {
		os.Exit(globalCmd.Ret)
	}
}
