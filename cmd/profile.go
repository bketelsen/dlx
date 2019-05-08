// Copyright © 2019 bketelsen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"devlx/path"

	"github.com/gobuffalo/packr/v2"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	w bool
)

var profiles = []string{"gui", "cli"}

var profileCmd = &cobra.Command{
	Use:   "profile [name]",
	Short: "create or replace the provisioning profile for devlx",
	Long: `Profile creates or replaces the 'gui' and 'cli' profiles in lxc that allow you
to connect to running containers and possibly display X11 applications on the host. Run with
no arguments to create or update all required profiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			args = profiles
		}

		if w {
			writeProfile(args)
		} else {
			listProfile(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	profileCmd.PersistentFlags().StringVar(&config.Network, "network", viper.GetString("network"), "the name of your network device e.g. 'enp5s0'")

	profileCmd.PersistentFlags().BoolVarP(&w, "write", "w", false, "Create or update a profile")
}

func initProfileTemplates() error {
	pbox := packr.New("profiles", "../templates/profiles")

	err := os.MkdirAll(filepath.Join(path.GetConfigPath(), "profiles"), 0755)
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

func fillProfileTemplate(p string) ([]byte, error) {

	filename := p + ".yaml"

	log.Running("Reading profile " + p)
	profTmpl, err := ioutil.ReadFile(filepath.Join(path.GetConfigPath(), "profiles", filename))
	if err != nil {
		log.Error("Unable to read profile tempalte" + err.Error())
	}

	profTmplS := string(profTmpl)

	log.Running("Parsing profile template " + p)
	tmpl, err := template.New("profile").Parse(profTmplS)
	if err != nil {
		log.Error("Error parsing profile template" + err.Error())
		return nil, err
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, config)
	if err != nil {
		log.Error("Error executing on profile template" + err.Error())
		return nil, err
	}

	parsedTmpl := buf.Bytes()

	return parsedTmpl, nil
}

func writeProfile(profiles []string) error {
	c, err := lxd.ConnectLXDUnix(config.LxdSocket, nil)
	if err != nil {
		log.Error("Unable to connect: " + err.Error())
		os.Exit(1)
	}

	for _, p := range profiles {
		exists := true
		_, etag, err := c.GetProfile(p)
		if err != nil {
			exists = false
		}

		profileYaml, err := fillProfileTemplate(p)
		if err != nil {
			log.Error("Unable to fill profile: " + err.Error())
			return err
		}

		if exists {
			log.Running("Updating profile " + p)
			var profile api.ProfilePut
			err = yaml.Unmarshal(profileYaml, &profile)
			if err != nil {
				log.Error("Parsing Profile YAML : " + err.Error())
				return err
			}
			err = c.UpdateProfile(p, profile, etag)
			if err != nil {
				log.Error("Create Profile : " + err.Error())
				return err
			}

			log.Success("Updating profile " + p)
		} else {

			log.Running("Creating profile " + p)
			var profile api.ProfilesPost
			err = yaml.Unmarshal(profileYaml, &profile)
			if err != nil {
				log.Error("Parsing Profile Yaml : " + err.Error())
				return err
			}
			profile.Name = p
			err = c.CreateProfile(profile)
			if err != nil {
				log.Error("Create Profile : " + err.Error())
				return err
			}
			log.Success("Creating profile " + p)
		}
	}
	return nil
}

func listProfile(profules []string) {
	c, err := lxd.ConnectLXDUnix(config.LxdSocket, nil)
	if err != nil {
		log.Error("Unable to connect: " + err.Error())
		os.Exit(1)
	}

	for _, p := range profiles {
		exists := true
		prof, _, err := c.GetProfile(p)
		if err != nil {
			exists = false
		}

		if exists {
			fmt.Println("[", p, "]: ", prof)
		} else {
			log.Error("Profile doesn't exist in LXD: " + p)
		}
	}
}
