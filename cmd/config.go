// Copyright © 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"fmt"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"devlx/path"

	"github.com/gobuffalo/packr/v2"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type devlxConfig struct {
	network   string
	display   string
	template  string
	image     string
	lxdSocket string
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage global configurations",
	Long: `Helps manage your devlx configuration
	Running bare 'devlx config' collects config info needed to configure devlx.

	Running coffig with one or more flags will write that value to the config.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Running("Setting configuration")

		flagSet := false

		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if flag.Changed {
				flagSet = true

				viper.Set(flag.Name, flag.Value.String())
			}
		})

		if flagSet == false {
			if err := initConfigFile(); err != nil {
				log.Error(`Error getting config values` + err.Error())
				os.Exit(1)

			}
		}

		if err := viper.WriteConfig(); err != nil {
			log.Error(`Error writing update to config file` + err.Error())
			os.Exit(1)

		}

		log.Success("Completed Configuration")
	},
}

func initConfigFile() error {

	if err := validateLxdSetup(); err != nil {
		return err
	}

	if err := determineLxdSocket(); err != nil {
		return err
	}

	if err := determineNetwork(); err != nil {
		return err
	}

	if err := determineTemplate(); err != nil {
		return err
	}

	if err := determineImage(); err != nil {
		return err
	}

	if err := determineDisplay(); err != nil {
		return err
	}

	return nil
}

func createTemplates() error {
	box := packr.New("provision", "../templates/provision")

	err := os.MkdirAll(filepath.Join(path.GetConfigPath(), "provision"), 0755)
	if err != nil {
		return err
	}
	for _, tpl := range box.List() {
		bb, err := box.Find(tpl)
		if err != nil {
			return err
		}
		f, err := os.Create(filepath.Join(path.GetConfigPath(), "provision", tpl))
		if err != nil {
			return err
		}
		_, err = f.Write([]byte(bb))
		if err != nil {
			return err
		}
	}

	pbox := packr.New("profiles", "../templates/profiles")
	err = os.MkdirAll(filepath.Join(path.GetConfigPath(), "profiles"), 0755)
	if err != nil {
		return err
	}
	for _, tpl := range pbox.List() {
		bb, err := pbox.Find(tpl)
		if err != nil {
			return err
		}
		f, err := os.Create(filepath.Join(path.GetConfigPath(), "profiles", tpl))
		if err != nil {
			return err
		}
		_, err = f.Write([]byte(bb))
		if err != nil {
			return err
		}
	}
	return nil

}

func createRelationsStore() error {
	// create config storage directory and file

	err := os.MkdirAll(filepath.Join(path.GetConfigPath(), "templates"), 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(path.GetConfigPath(), "templates", "relations.yaml"))
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}

func determineDisplay() error {
	return nil
	//TODO
}

func determineImage() error {
	prompt := promptui.Select{
		Label: "Select default Ubuntu OS Image",
		Items: []string{"19.04", "18.10", "18.04", "16.04", "14.04"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	viper.Set("image", result)
	return nil
}

func determineLxdSocket() error {
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
		log.Error(`No LXD Socket found, are you sure LXD is installed?`)
		os.Exit(1)
	}

	viper.Set("lxd-socket", lxdSocket)

	return nil
}

func determineNetwork() error {
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

func determineTemplate() error {

	prompt := promptui.Select{
		Label: "Select default template",
		Items: []string{"gui", "cli"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	viper.Set("template", result)
	return nil
}

func validateLxdSetup() error {
	curUser, err := user.Current()
	if err != nil {
		log.Error("Unable to find default UID from OS.")
		return err
	}

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

	return nil
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVar(&config.network, "network", viper.GetString("network"), "Set default Network interface for bridging")
	configCmd.Flags().StringVar(&config.display, "display", viper.GetString("display"), "Set default Display name for gui application windows")
	configCmd.Flags().StringVar(&config.template, "template", viper.GetString("template"), "Set default template for creating templates")
	configCmd.Flags().StringVar(&config.image, "image", viper.GetString("image"), "Set default  Ubuntu image for creating templates")

}
