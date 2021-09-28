package lxc

import (
	"os"

	"github.com/bketelsen/dlx/state"
	cli "github.com/lxc/lxd/shared/cmd"
	"github.com/lxc/lxd/shared/i18n"
	"github.com/lxc/lxd/shared/version"
	"github.com/spf13/cobra"
)

type CmdLxc struct {
	Global *state.Global
}

func (c *CmdLxc) Command() *cobra.Command {
	lxc := &cobra.Command{}
	lxc.Use = "lxc"
	lxc.Short = i18n.G("Command line client for LXD")
	lxc.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Command line client for LXD

All of LXD's features can be driven through the various commands below.
For help with any of those, simply call them with --help.`))
	lxc.SilenceUsage = true
	lxc.SilenceErrors = true
	lxc.CompletionOptions = cobra.CompletionOptions{DisableDefaultCmd: true}

	// Wrappers
	lxc.PersistentPreRunE = c.Global.PreRun
	lxc.PersistentPostRunE = c.Global.PostRun

	// Version handling
	lxc.SetVersionTemplate("{{.Version}}\n")
	lxc.Version = version.Version

	// alias sub-command
	aliasCmd := cmdAlias{global: c.Global}
	lxc.AddCommand(aliasCmd.Command())

	// cluster sub-command
	clusterCmd := cmdCluster{global: c.Global}
	lxc.AddCommand(clusterCmd.Command())

	// config sub-command
	configCmd := cmdConfig{global: c.Global}
	lxc.AddCommand(configCmd.Command())

	// console sub-command
	consoleCmd := cmdConsole{global: c.Global}
	lxc.AddCommand(consoleCmd.Command())

	// copy sub-command
	copyCmd := cmdCopy{global: c.Global}
	lxc.AddCommand(copyCmd.Command())

	// delete sub-command
	deleteCmd := cmdDelete{global: c.Global}
	lxc.AddCommand(deleteCmd.Command())

	// exec sub-command
	execCmd := cmdExec{global: c.Global}
	lxc.AddCommand(execCmd.Command())

	// export sub-command
	exportCmd := cmdExport{global: c.Global}
	lxc.AddCommand(exportCmd.Command())

	// file sub-command
	fileCmd := cmdFile{global: c.Global}
	lxc.AddCommand(fileCmd.Command())

	// import sub-command
	importCmd := cmdImport{global: c.Global}
	lxc.AddCommand(importCmd.Command())

	// info sub-command
	infoCmd := cmdInfo{global: c.Global}
	lxc.AddCommand(infoCmd.Command())

	// image sub-command
	imageCmd := cmdImage{global: c.Global}
	lxc.AddCommand(imageCmd.Command())

	// init sub-command
	initCmd := cmdInit{global: c.Global}
	lxc.AddCommand(initCmd.Command())

	// launch sub-command
	launchCmd := cmdLaunch{global: c.Global, init: &initCmd}
	lxc.AddCommand(launchCmd.Command())

	// list sub-command
	listCmd := cmdList{global: c.Global}
	lxc.AddCommand(listCmd.Command())

	// manpage sub-command
	manpageCmd := cmdManpage{global: c.Global}
	lxc.AddCommand(manpageCmd.Command())

	// monitor sub-command
	monitorCmd := cmdMonitor{global: c.Global}
	lxc.AddCommand(monitorCmd.Command())

	// move sub-command
	moveCmd := cmdMove{global: c.Global}
	lxc.AddCommand(moveCmd.Command())

	// network sub-command
	networkCmd := cmdNetwork{global: c.Global}
	lxc.AddCommand(networkCmd.Command())

	// operation sub-command
	operationCmd := cmdOperation{global: c.Global}
	lxc.AddCommand(operationCmd.Command())

	// pause sub-command
	pauseCmd := cmdPause{global: c.Global}
	lxc.AddCommand(pauseCmd.Command())

	// publish sub-command
	publishCmd := cmdPublish{global: c.Global}
	lxc.AddCommand(publishCmd.Command())

	// profile sub-command
	profileCmd := cmdProfile{global: c.Global}
	lxc.AddCommand(profileCmd.Command())

	// profile sub-command
	projectCmd := cmdProject{global: c.Global}
	lxc.AddCommand(projectCmd.Command())

	// query sub-command
	queryCmd := cmdQuery{global: c.Global}
	lxc.AddCommand(queryCmd.Command())

	// rename sub-command
	renameCmd := cmdRename{global: c.Global}
	lxc.AddCommand(renameCmd.Command())

	// restart sub-command
	restartCmd := cmdRestart{global: c.Global}
	lxc.AddCommand(restartCmd.Command())

	// remote sub-command
	remoteCmd := cmdRemote{global: c.Global}
	lxc.AddCommand(remoteCmd.Command())

	// restore sub-command
	restoreCmd := cmdRestore{global: c.Global}
	lxc.AddCommand(restoreCmd.Command())

	// snapshot sub-command
	snapshotCmd := cmdSnapshot{global: c.Global}
	lxc.AddCommand(snapshotCmd.Command())

	// storage sub-command
	storageCmd := cmdStorage{global: c.Global}
	lxc.AddCommand(storageCmd.Command())

	// start sub-command
	startCmd := cmdStart{global: c.Global}
	lxc.AddCommand(startCmd.Command())

	// stop sub-command
	stopCmd := cmdStop{global: c.Global}
	lxc.AddCommand(stopCmd.Command())

	// version sub-command
	versionCmd := cmdVersion{global: c.Global}
	lxc.AddCommand(versionCmd.Command())

	// warning sub-command
	warningCmd := cmdWarning{global: c.Global}
	lxc.AddCommand(warningCmd.Command())

	// Get help command
	lxc.InitDefaultHelpCmd()
	var help *cobra.Command
	for _, cmd := range lxc.Commands() {
		if cmd.Name() == "help" {
			help = cmd
			break
		}
	}

	// Help flags
	lxc.Flags().BoolVar(&c.Global.FlagHelpAll, "all", false, i18n.G("Show less common commands"))
	help.Flags().BoolVar(&c.Global.FlagHelpAll, "all", false, i18n.G("Show less common commands"))

	// Deal with --all flag
	err := lxc.ParseFlags(os.Args[1:])
	if err == nil {
		if c.Global.FlagHelpAll {
			// Show all commands
			for _, cmd := range lxc.Commands() {
				cmd.Hidden = false
			}
		}
	}

	return lxc
}
