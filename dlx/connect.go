package dlx

import (
	"os"
	"strconv"
	"strings"

	"github.com/bketelsen/dlx/dlx/config"
	"github.com/bketelsen/dlx/state"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	cli "github.com/lxc/lxd/shared/cmd"
	"github.com/lxc/lxd/shared/i18n"
	"github.com/lxc/lxd/shared/logger"
	"github.com/lxc/lxd/shared/termios"
)

type CmdConnect struct {
	Global *state.Global

	flagEnvironment []string
}

func (c *CmdConnect) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = usage("connect", i18n.G("<instance> [flags]"))
	cmd.Short = i18n.G("Connect to a login shell of an instance")
	cmd.Long = cli.FormatSection(i18n.G("Description"), i18n.G(
		`Connect to a login shell of an instance.`))

	cmd.RunE = c.Run
	cmd.Flags().StringArrayVar(&c.flagEnvironment, "env", nil, i18n.G("Environment variable to set (e.g. HOME=/home/foo)")+"``")

	return cmd
}

func (c *CmdConnect) sendTermSize(control *websocket.Conn) error {
	width, height, err := termios.GetSize(getStdoutFd())
	if err != nil {
		return err
	}

	logger.Debugf("Window size is now: %dx%d", width, height)

	msg := api.InstanceExecControl{}
	msg.Command = "window-resize"
	msg.Args = make(map[string]string)
	msg.Args["width"] = strconv.Itoa(width)
	msg.Args["height"] = strconv.Itoa(height)

	return control.WriteJSON(msg)
}

func (c *CmdConnect) Run(cmd *cobra.Command, args []string) error {
	conf := c.Global.Conf
	cfg, err := config.Get()
	if err != nil {
		return errors.Wrap(err, "unable to get configuration")
	}

	// Quick checks.
	exit, err := c.Global.CheckArgs(cmd, args, 1, -1)
	if exit {
		return err
	}

	// Connect to the daemon
	remote, name, err := conf.ParseRemote(args[0])
	if err != nil {
		return err
	}

	d, err := conf.GetInstanceServer(remote)
	if err != nil {
		return err
	}

	// Set the environment
	env := map[string]string{}
	myTerm, ok := c.getTERM()
	if ok {
		env["TERM"] = myTerm
	}

	for _, arg := range c.flagEnvironment {
		pieces := strings.SplitN(arg, "=", 2)
		value := ""
		if len(pieces) > 1 {
			value = pieces[1]
		}
		env[pieces[0]] = value
	}

	// Configure the terminal
	stdinFd := getStdinFd()
	stdoutFd := getStdoutFd()

	stdinTerminal := termios.IsTerminal(stdinFd)
	stdoutTerminal := termios.IsTerminal(stdoutFd)

	// Determine interaction mode
	interactive := true

	// Record terminal state
	var oldttystate *termios.State
	if interactive && stdinTerminal {
		oldttystate, err = termios.MakeRaw(stdinFd)
		if err != nil {
			return err
		}

		defer termios.Restore(stdinFd, oldttystate)
	}

	// Setup interactive console handler
	handler := c.controlSocketHandler
	if !interactive {
		handler = nil
	}

	// Grab current terminal dimensions
	var width, height int
	if stdoutTerminal {
		width, height, err = termios.GetSize(getStdoutFd())
		if err != nil {
			return err
		}
	}

	stdin := os.Stdin

	stdout := getStdout()

	// Prepare the command
	user := cfg.Remotes[c.Global.Conf.DefaultRemote].User
	//cmdargs := args[1:]

	req := api.InstanceExecPost{
		Command:     []string{"sudo", "-u", user, "/bin/bash", "--login"},
		WaitForWS:   true,
		Interactive: interactive,
		Environment: env,
		Width:       width,
		Height:      height,
		Cwd:         "/home/" + user,
	}

	execArgs := lxd.InstanceExecArgs{
		Stdin:    stdin,
		Stdout:   stdout,
		Stderr:   os.Stderr,
		Control:  handler,
		DataDone: make(chan bool),
	}

	// Run the command in the instance
	op, err := d.ExecInstance(name, req, &execArgs)
	if err != nil {
		return err
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return err
	}
	opAPI := op.Get()

	// Wait for any remaining I/O to be flushed
	<-execArgs.DataDone

	c.Global.Ret = int(opAPI.Metadata["return"].(float64))
	return nil
}
