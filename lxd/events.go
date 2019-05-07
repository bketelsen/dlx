// Copyright © 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package lxd

import "time"

type ConnectionCreated struct {
	conn      *Client
	StartTime time.Time
}

func (e ConnectionCreated) Name() string {
	return e.conn.URL
}
func (e ConnectionCreated) Created() time.Time {
	return e.StartTime
}

func NewConnectionCreated(conn *Client) *ConnectionCreated {
	e := &ConnectionCreated{
		conn:      conn,
		StartTime: time.Now(),
	}
	return e
}

type State string

const (
	Creating     State = "creating"
	Created      State = "created"
	Starting     State = "starting"
	Started      State = "started"
	Stopped      State = "stopped"
	Stopping     State = "stopping"
	Completed    State = "completed"
	Removing     State = "removing"
	Removed      State = "removed"
	Provisioning State = "provisioning"
	Provisioned  State = "provisioned"
)

type ContainerState struct {
	ContainerState State
	ContainerName  string
	StartTime      time.Time
}

func (e ContainerState) Name() string {
	return e.ContainerName + "\t" + string(e.ContainerState)
}
func (e ContainerState) Created() time.Time {
	return e.StartTime
}

func NewContainerState(name string, state State) *ContainerState {
	e := &ContainerState{
		ContainerState: state,
		ContainerName:  name,
		StartTime:      time.Now(),
	}
	return e
}

type ExecState struct {
	CommandState  State
	Command       string
	ContainerName string
	StartTime     time.Time
}

func (e ExecState) Name() string {
	return e.ContainerName + "\t" + string(e.CommandState) + "\t" + e.Command
}
func (e ExecState) Created() time.Time {
	return e.StartTime
}

func NewExecState(name string, command string, state State) *ExecState {
	e := &ExecState{
		CommandState:  state,
		Command:       command,
		ContainerName: name,
		StartTime:     time.Now(),
	}
	return e
}

type CopyState struct {
	OperationState State
	File           string
	ContainerName  string
	StartTime      time.Time
}

func (e CopyState) Name() string {
	return e.ContainerName + "\t" + string(e.OperationState) + "\t" + e.File
}
func (e CopyState) Created() time.Time {
	return e.StartTime
}

func NewCopyState(name string, file string, state State) *CopyState {
	e := &CopyState{
		OperationState: state,
		File:           file,
		ContainerName:  name,
		StartTime:      time.Now(),
	}
	return e
}
