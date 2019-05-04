package config

import (
	"devlx/path"
	"errors"
	"fmt"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/dixonwille/wlog"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

var log wlog.UI

type GlobalConfig struct {
	Network   string
	LxdSocket string
	Uid       string
	Display   string
}

// initConfig reads in config file and ENV variables if set.
func (c GlobalConfig) New(verbose bool, cfgFile string) GlobalConfig {

	return c
}

func (c *GlobalConfig) GetConfig() GlobalConfig {
	return *c
}

func (c *GlobalConfig) WriteConfig() error {
	err := viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func getConfigValues() error {

	//Sets default to curent user's uid if $UID is not set in env. (This should be on most linux systems)
	curUser, err := user.Current()
	if err != nil {
		log.Error("Unable to find default UID from OS.")
		return err
	}

	viper.SetDefault("uid", curUser.Uid) // probably don't need to store this in global config?

	//confirm user is in lxd group
	lxdGroup, err := user.LookupGroup("lxd")
	if err != nil {
		log.Error("Unable to get lxd group Id from OS, are you sure lxd is installed?")
		return err
	}

	userGroups, err := curUser.GroupIds()
	if err != nil {
		log.Error("Unable to get user's group IDs from OS.")
		return err
	}

	userInGroup := false

	for _, gid := range userGroups {
		if gid == lxdGroup.Gid {
			userInGroup = true
		}
	}

	if !userInGroup {
		log.Error(fmt.Sprintf("The current user %s is not in the 'lxd' group. Please add the user by running 'adduser %s lxd' in terminal then logging out and back in before rerunning init.", curUser.Name, curUser.Username))
	}

	//Find and set LXD Unix socket.
	lxdSocket := ""
	possibleLxdSockets := []string{
		"/var/snap/lxd/common/lxd/unix.socket",
		"/var/lib/lxd/unix.socket",
	}

	for _, socket := range possibleLxdSockets {
		if _, err := os.Stat(socket); err == nil {
			lxdSocket = socket
			break
		}
	}

	if lxdSocket == "" {
		log.Error("LXD Unix socket not found, are you sure LXD is installed?")
		os.Exit(1)
	}

	viper.Set("lxdSocket", lxdSocket)

	//Get Network adapters
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	var interfaceNames []string

	for _, inter := range interfaces {

		if strings.Contains(inter.Name, "lo") || strings.Contains(inter.Name, "tun") || strings.Contains(inter.Name, "docker") {
			continue
		}

		interfaceNames = append(interfaceNames, inter.Name)
	}

	if l := len(interfaceNames); l > 1 {
		prompt := promptui.Select{
			Label: "Select Network Adapter",
			Items: interfaceNames,
		}

		_, result, err := prompt.Run()

		if err != nil {
			return err
		}

		viper.Set("network", result)
	} else if l == 1 {
		viper.Set("network", interfaceNames[0])
	} else {
		//in the future we should probably create a host network like docker does.
		return errors.New("No network interfaces available")
	}

	return nil

}

func createConfig() error {
	//make config directory and file
	err := os.MkdirAll(filepath.Join(path.GetConfigPath()), 0755)

	// f, err := os.Create(filepath.Join(path.GetConfigPath(), "devlx.yaml"))
	err = viper.WriteConfig()
	if err != nil {
		return err
	}
	// _, err = f.Write([]byte(configTemplate))
	// if err != nil {
	// 	return err
	// }
	return nil
}

func checkConfig() error {
	_, err := os.Stat(filepath.Join(path.GetConfigPath(), "devlx.yaml"))
	if err != nil {
		return err
	}
	return nil
}
