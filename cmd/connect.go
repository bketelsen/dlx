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
	"log"
	"os"
	"syscall"

	"github.com/buger/goterm"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/lxc/lxd/shared/termios"
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to a running container",
	Long:  `Connect to a running container.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name = args[0]
		// Connect to LXD over the Unix socket
		c, err := lxd.ConnectLXDUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
		if err != nil {
			log.Fatal("Connect:", err)
		}
		terminalHeight := goterm.Height()
		terminalWidth := goterm.Width()
		// Setup the exec request
		environ := make(map[string]string)
		environ["TERM"] = os.Getenv("TERM")
		req := api.ContainerExecPost{
			Command:     []string{"/bin/bash", "-c", "sudo --user ubuntu --login"},
			WaitForWS:   true,
			Interactive: true,
			Width:       terminalWidth,
			Height:      terminalHeight,
			Environment: environ,
		}

		// Setup the exec arguments (fds)
		largs := lxd.ContainerExecArgs{
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		// Setup the terminal (set to raw mode)
		if req.Interactive {
			cfd := int(syscall.Stdin)
			oldttystate, err := termios.MakeRaw(cfd)
			if err != nil {
				log.Fatal("Make Raw Terminal", err)
			}

			defer termios.Restore(cfd, oldttystate)
		}

		// Get the current state
		op, err := c.ExecContainer(name, req, &largs)
		if err != nil {
			log.Fatal("Exec:", err)
		}

		// Wait for it to complete
		err = op.Wait()
		if err != nil {
			log.Fatal("Wait:", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
