//go:build !windows
// +build !windows

package dlx

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/gorilla/websocket"
	"golang.org/x/sys/unix"

	"github.com/lxc/lxd/shared/logger"
)

func (c *CmdConsole) controlSocketHandler(control *websocket.Conn) {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, unix.SIGWINCH)

	for {
		sig := <-ch
		logger.Debugf("Received '%s signal', updating window geometry.", sig)
		err := c.sendTermSize(control)
		if err != nil {
			logger.Debugf("error setting term size %s", err)
			break
		}
	}

	closeMsg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	control.WriteMessage(websocket.CloseMessage, closeMsg)
}

func (c *CmdConsole) findCommand(name string) string {
	path, _ := exec.LookPath(name)
	return path
}
