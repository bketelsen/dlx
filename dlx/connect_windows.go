//go:build windows
// +build windows

package dlx

import (
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"golang.org/x/sys/windows"

	"github.com/lxc/lxd/shared/api"
	"github.com/lxc/lxd/shared/logger"
)

func (c *CmdConnect) getTERM() (string, bool) {
	return "dumb", true
}

func (c *CmdConnect) controlSocketHandler(control *websocket.Conn) {
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, os.Interrupt)

	closeMsg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	defer control.WriteMessage(websocket.CloseMessage, closeMsg)

	for {
		sig := <-ch
		switch sig {
		case os.Interrupt:
			logger.Debugf("Received '%s signal', forwarding to executing program.", sig)
			err := c.forwardSignal(control, windows.SIGINT)
			if err != nil {
				logger.Debugf("Failed to forward signal '%s'.", windows.SIGINT)
				return
			}
		default:
			break
		}
	}
}

func (c *CmdConnect) forwardSignal(control *websocket.Conn, sig windows.Signal) error {
	logger.Debugf("Forwarding signal: %s", sig)

	msg := api.InstanceExecControl{}
	msg.Command = "signal"
	msg.Signal = int(sig)

	return control.WriteJSON(msg)
}
