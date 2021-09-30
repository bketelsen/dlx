package dlx

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/bketelsen/dlx/config"
	"github.com/bketelsen/dlx/state"
	"github.com/lxc/lxd/lxc/utils"
	"github.com/lxc/lxd/shared"
	"github.com/lxc/lxd/shared/api"
	cli "github.com/lxc/lxd/shared/cmd"
	"github.com/lxc/lxd/shared/i18n"
)

// Start
type CmdStart struct {
	Global *state.Global
	action *CmdAction
}

func (c *CmdStart) Command() *cobra.Command {
	cmdAction := CmdAction{Global: c.Global}
	c.action = &cmdAction

	cmd := c.action.Command("start")
	cmd.Use = usage("start", i18n.G("[<remote>:]<instance> [[<remote>:]<instance>...]"))
	cmd.Short = i18n.G("Start instances")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Start instances`))

	return cmd
}

// Pause
type CmdPause struct {
	Global *state.Global
	action *CmdAction
}

func (c *CmdPause) Command() *cobra.Command {
	cmdAction := CmdAction{Global: c.Global}
	c.action = &cmdAction

	cmd := c.action.Command("pause")
	cmd.Use = usage("pause", i18n.G("[<remote>:]<instance> [[<remote>:]<instance>...]"))
	cmd.Short = i18n.G("Pause instances")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Pause instances`))
	cmd.Hidden = true

	return cmd
}

// Restart
type CmdRestart struct {
	Global *state.Global
	action *CmdAction
}

func (c *CmdRestart) Command() *cobra.Command {
	cmdAction := CmdAction{Global: c.Global}
	c.action = &cmdAction

	cmd := c.action.Command("restart")
	cmd.Use = usage("restart", i18n.G("[<remote>:]<instance> [[<remote>:]<instance>...]"))
	cmd.Short = i18n.G("Restart instances")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Restart instances

The opposite of "lxc pause" is "lxc start".`))

	return cmd
}

// Stop
type CmdStop struct {
	Global *state.Global
	action *CmdAction
}

func (c *CmdStop) Command() *cobra.Command {
	cmdAction := CmdAction{Global: c.Global}
	c.action = &cmdAction

	cmd := c.action.Command("stop")
	cmd.Use = usage("stop", i18n.G("[<remote>:]<instance> [[<remote>:]<instance>...]"))
	cmd.Short = i18n.G("Stop instances")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Stop instances`))

	return cmd
}

type CmdAction struct {
	Global *state.Global

	flagAll       bool
	flagConsole   string
	flagForce     bool
	flagStateful  bool
	flagStateless bool
	flagTimeout   int
}

func (c *CmdAction) Command(action string) *cobra.Command {
	cmd := &cobra.Command{}
	cmd.RunE = c.Run

	cmd.Flags().BoolVar(&c.flagAll, "all", false, i18n.G("Run against all instances"))

	if action == "stop" {
		cmd.Flags().BoolVar(&c.flagStateful, "stateful", false, i18n.G("Store the instance state"))
	} else if action == "start" {
		cmd.Flags().BoolVar(&c.flagStateless, "stateless", false, i18n.G("Ignore the instance state"))
	}

	if shared.StringInSlice(action, []string{"start", "restart", "stop"}) {
		cmd.Flags().StringVar(&c.flagConsole, "console", "", i18n.G("Immediately attach to the console")+"``")
		cmd.Flags().Lookup("console").NoOptDefVal = "console"
	}

	if shared.StringInSlice(action, []string{"restart", "stop"}) {
		cmd.Flags().BoolVarP(&c.flagForce, "force", "f", false, i18n.G("Force the instance to shutdown"))
		cmd.Flags().IntVar(&c.flagTimeout, "timeout", -1, i18n.G("Time to wait for the instance before killing it")+"``")
	}

	return cmd
}

func (c *CmdAction) doActionAll(action string, resource state.RemoteResource) error {
	if resource.Name != "" {
		// both --all and instance name given.
		return fmt.Errorf(i18n.G("Both --all and instance name given"))
	}

	remote := resource.Remote
	d, err := c.Global.Conf.GetInstanceServer(remote)
	if err != nil {
		return err
	}

	// Pause is called freeze.
	if action == "pause" {
		action = "freeze"
	}

	// Only store state if asked to.
	state := false
	if action == "stop" && c.flagStateful {
		state = true
	}

	req := api.InstancesPut{
		State: &api.InstanceStatePut{
			Action:   action,
			Timeout:  c.flagTimeout,
			Force:    c.flagForce,
			Stateful: state,
		},
	}

	// Update all instances.
	op, err := d.UpdateInstances(req, "")
	if err != nil {
		return err
	}

	progress := utils.ProgressRenderer{
		Quiet: c.Global.FlagQuiet,
	}

	_, err = op.AddHandler(progress.UpdateOp)
	if err != nil {
		progress.Done("")
		return err
	}

	err = utils.CancelableWait(op, &progress)
	if err != nil {
		progress.Done("")
		return err
	}

	progress.Done("")

	return nil
}

func (c *CmdAction) doAction(action string, conf *config.Config, nameArg string) error {
	state := false

	// Pause is called freeze
	if action == "pause" {
		action = "freeze"
	}

	// Only store state if asked to
	if action == "stop" && c.flagStateful {
		state = true
	}

	if action == "stop" && c.flagForce && c.flagConsole != "" {
		return fmt.Errorf(i18n.G("--console can't be used while forcing instance shutdown"))
	}

	remote, name, err := conf.ParseRemote(nameArg)
	if err != nil {
		return err
	}

	d, err := conf.GetInstanceServer(remote)
	if err != nil {
		return err
	}

	if name == "" {
		return fmt.Errorf(i18n.G("Must supply instance name for: ")+"\"%s\"", nameArg)
	}

	if action == "start" {
		current, _, err := d.GetInstance(name)
		if err != nil {
			return err
		}

		// "start" for a frozen instance means "unfreeze"
		if current.StatusCode == api.Frozen {
			action = "unfreeze"
		}

		// Always restore state (if present) unless asked not to
		if action == "start" && current.Stateful && !c.flagStateless {
			state = true
		}
	}

	req := api.InstanceStatePut{
		Action:   action,
		Timeout:  c.flagTimeout,
		Force:    c.flagForce,
		Stateful: state,
	}

	op, err := d.UpdateInstanceState(name, req, "")
	if err != nil {
		return err
	}

	if action == "stop" && c.flagConsole != "" {
		// Handle console attach
		console := CmdConsole{}
		console.Global = c.Global
		console.flagType = c.flagConsole
		return console.Console(d, name)
	}

	progress := utils.ProgressRenderer{
		Quiet: c.Global.FlagQuiet,
	}
	_, err = op.AddHandler(progress.UpdateOp)
	if err != nil {
		progress.Done("")
		return err
	}

	// Wait for operation to finish
	err = utils.CancelableWait(op, &progress)
	if err != nil {
		progress.Done("")
		return fmt.Errorf("%s\n"+i18n.G("Try `lxc info --show-log %s` for more info"), err, nameArg)
	}

	progress.Done("")

	// Handle console attach
	if c.flagConsole != "" {
		console := CmdConsole{}
		console.Global = c.Global
		console.flagType = c.flagConsole
		return console.Console(d, name)
	}

	return nil
}

func (c *CmdAction) Run(cmd *cobra.Command, args []string) error {
	conf := c.Global.Conf

	var names []string
	if c.flagAll {
		// If no server passed, use current default.
		if len(args) == 0 {
			args = []string{fmt.Sprintf("%s:", conf.DefaultRemote)}
		}

		// Get all the servers.
		resources, err := c.Global.ParseServers(args...)
		if err != nil {
			return err
		}

		for _, resource := range resources {
			// We don't allow instance names with --all.
			if resource.Name != "" {
				return fmt.Errorf(i18n.G("Both --all and instance name given"))
			}

			// See if we can use the bulk API.
			if resource.Server.HasExtension("instance_bulk_state_change") {
				err = c.doActionAll(cmd.Name(), resource)
				if err != nil {
					return fmt.Errorf("%s: %v", resource.Remote, err)
				}

				continue
			}

			ctslist, err := resource.Server.GetInstances(api.InstanceTypeAny)
			if err != nil {
				return err
			}

			for _, ct := range ctslist {
				switch cmd.Name() {
				case "start":
					if ct.StatusCode == api.Running {
						continue
					}
				case "stop":
					if ct.StatusCode == api.Stopped {
						continue
					}
				}
				names = append(names, fmt.Sprintf("%s:%s", resource.Remote, ct.Name))
			}
		}
	} else {
		names = args

		if len(args) == 0 {
			cmd.Usage()
			return nil
		}
	}

	if c.flagConsole != "" {
		if c.flagAll {
			return fmt.Errorf(i18n.G("--console can't be used with --all"))
		}

		if len(names) != 1 {
			return fmt.Errorf(i18n.G("--console only works with a single instance"))
		}
	}

	// Run the action for every listed instance
	results := runBatch(names, func(name string) error { return c.doAction(cmd.Name(), conf, name) })

	// Single instance is easy
	if len(results) == 1 {
		return results[0].err
	}

	// Do fancier rendering for batches
	success := true

	for _, result := range results {
		if result.err == nil {
			continue
		}

		success = false
		msg := fmt.Sprintf(i18n.G("error: %v"), result.err)
		for _, line := range strings.Split(msg, "\n") {
			fmt.Fprintf(os.Stderr, "%s: %s\n", result.name, line)
		}
	}

	if !success {
		fmt.Fprintln(os.Stderr, "")
		return fmt.Errorf(i18n.G("Some instances failed to %s"), cmd.Name())
	}

	return nil
}
