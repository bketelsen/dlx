// Copyright (c) 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package dlx

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/bketelsen/dlx/state"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/lxc/lxd/shared/i18n"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type CmdMonitor struct {
	Cmd             *cobra.Command
	Global          *state.Global
	flagType        []string
	flagAllProjects bool
}

func (c *CmdMonitor) Command() *cobra.Command {

	// removeCmd represents the remove command
	var removeCmd = &cobra.Command{
		Use:   "monitor",
		Short: "monitor lxd events",
		Long: `Monitor watches lxd events. It is meant to be run continuously
on the LXD server to provision instances as they are created.`,
		RunE: c.Run,
	}

	removeCmd.Flags().BoolVar(&c.flagAllProjects, "all-projects", false, "Show events from all projects")
	removeCmd.Flags().StringArrayVar(&c.flagType, "type", nil, "Event type to listen for"+"``")

	return removeCmd
}
func (c *CmdMonitor) Run(cmd *cobra.Command, args []string) error {
	conf := c.Global.Conf

	var err error
	var remote string

	// Quick checks.
	exit, err := c.Global.CheckArgs(cmd, args, 0, 1)
	if exit {
		return err
	}

	// Connect to the event source.
	if len(args) == 0 {
		remote, _, err = conf.ParseRemote("")
		if err != nil {
			return err
		}
	} else {
		remote, _, err = conf.ParseRemote(args[0])
		if err != nil {
			return err
		}
	}

	d, err := conf.GetInstanceServer(remote)
	if err != nil {
		return err
	}

	if c.flagAllProjects {
		d = d.UseProject("*")
	}

	listener, err := d.GetEvents()
	if err != nil {
		return err
	}

	chError := make(chan error, 1)

	handler := func(event api.Event) {
		// Render as JSON (to expand RawMessage)
		jsonRender, err := json.Marshal(&event)
		if err != nil {
			chError <- err
			return
		}

		// Read back to a clean interface
		var rawEvent interface{}
		err = json.Unmarshal(jsonRender, &rawEvent)
		if err != nil {
			chError <- err
			return
		}

		// And now print the result.
		var render []byte

		render, err = json.Marshal(&rawEvent)
		if err != nil {
			chError <- err
			return
		}

		fmt.Printf("%s\n\n", render)
		fmt.Println(c.Global.FlagProject)
		go c.processEvent(event, d)
	}

	_, err = listener.AddHandler(c.flagType, handler)
	if err != nil {
		return err
	}

	go func() {
		chError <- listener.Wait()
	}()

	return <-chError

}

func (c *CmdMonitor) processEvent(event api.Event, d lxd.InstanceServer) {
	if event.Type == "lifecycle" {
		e := &api.EventLifecycle{}
		err := json.Unmarshal(event.Metadata, &e)
		if err != nil {
			log.Error(err.Error())
		}
		if e.Action == "instance-started" {
			// now do the stuff
			project := state.GetProject(c.Global.FlagProject)

			// ensure the host has the mount paths for project file storage
			err = project.CreateMountPath()
			if err != nil {
				log.Error(errors.Wrap(err, "creating mount path on host").Error())
				return
			}
			err = project.CreateCommonMountPath()
			if err != nil {
				log.Error(errors.Wrap(err, "creating common mount path on host").Error())
				return
			}
			name := filepath.Base(e.Source)

			log.Running("Checking if cloud-init exists")
			// Mount the project directory into container FS

			_, _, err := fileGetWrapper(d, name, "/usr/bin/cloud-init")
			if err == nil {

				// wait for cloud-init
				log.Running("Waiting for cloud-init to finish: " + name)
				time.Sleep(time.Second * 5)
				maxwait := 300
				sleeptime := 3
				for i := 0; i < maxwait; i += sleeptime {

					buf, resp, err := fileGetWrapper(d, name, "/var/lib/cloud/instance/boot-finished")
					if err == nil {
						fmt.Println(resp)
						contents, err := ioutil.ReadAll(buf)
						if err != nil {
							log.Error(errors.Wrap(err, "reading boot-finished").Error())
							return
						}
						fmt.Println(string(contents))
						break
					}
					log.Info(".")
					time.Sleep(time.Second * time.Duration(sleeptime))
				}
				log.Success("cloud-init finished")
			}

			devname := "persist"
			devSource := "source=" + project.InstanceMountPath(name)
			devPath := "path=" + project.ContainerMountPath()
			log.Info(devSource)
			log.Info(devPath)
			err = project.CreateInstanceMountPath(name)
			if err != nil {
				log.Error(errors.Wrap(err, "failed to create host mount path").Error())
				return
			}
			err = addDevice(d, name, []string{devname, "disk", devSource, devPath})
			if err != nil {
				log.Error(errors.Wrap(err, "failed to mount project directory").Error())
				return
			}

			// Mount the common directory into container FS
			cdevname := "common"
			cdevSource := "source=" + project.CommonMountPath()
			cdevPath := "path=" + project.ContainerCommonMountPath()

			err = addDevice(d, name, []string{cdevname, "disk", cdevSource, cdevPath})
			if err != nil {
				log.Error(errors.Wrap(err, "failed to mount project common storage directory").Error())
				return
			}

		}
	}

}
func addDevice(d lxd.InstanceServer, name string, args []string) error {

	// Add the device
	devname := args[0]
	device := map[string]string{}
	device["type"] = args[1]
	if len(args) > 2 {
		for _, prop := range args[2:] {
			results := strings.SplitN(prop, "=", 2)
			if len(results) != 2 {
				return fmt.Errorf("No value found in %q", prop)
			}
			k := results[0]
			v := results[1]
			device[k] = v
		}
	}

	inst, etag, err := d.GetInstance(name)
	if err != nil {
		return err
	}

	_, ok := inst.Devices[devname]
	if ok {
		return fmt.Errorf("The device already exists")
	}

	inst.Devices[devname] = device

	op, err := d.UpdateInstance(name, inst.Writable(), etag)
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}

	return nil
}

func fileGetWrapper(server lxd.InstanceServer, inst string, path string) (buf io.ReadCloser, resp *lxd.InstanceFileResponse, err error) {
	// Signal handling
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	// Operation handling
	chDone := make(chan bool)
	go func() {
		buf, resp, err = server.GetInstanceFile(inst, path)
		close(chDone)
	}()

	count := 0
	for {
		select {
		case <-chDone:
			return buf, resp, err
		case <-chSignal:
			count++

			if count == 3 {
				return nil, nil, fmt.Errorf(i18n.G("User signaled us three times, exiting. The remote operation will keep running"))
			}

			fmt.Println(i18n.G("Early server side processing of file transfer requests cannot be canceled (interrupt two more times to force)"))
		}
	}
}
