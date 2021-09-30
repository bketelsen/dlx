//go:build windows
// +build windows

package dlx

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gorilla/websocket"
)

func (c *CmdConsole) getTERM() (string, bool) {
	return "dumb", true
}

func (c *CmdConsole) controlSocketHandler(control *websocket.Conn) {
	// TODO: figure out what the equivalent of signal.SIGWINCH is on
	// windows and use that; for now if you resize your terminal it just
	// won't work quite correctly.
	err := c.sendTermSize(control)
	if err != nil {
		fmt.Printf("error setting term size %s\n", err)
	}
}

func (c *CmdConsole) findCommand(name string) string {
	path, _ := exec.LookPath(name)
	if path == "" {
		// Let's see if it's not in the usual location.
		programs, err := ioutil.ReadDir("\\Program Files")
		if err != nil {
			return ""
		}

		for _, entry := range programs {
			if strings.HasPrefix(entry.Name(), "VirtViewer") {
				return filepath.Join("\\Program Files", entry.Name(), "bin", "remote-viewer.exe")
			}
		}
	}

	return path
}
