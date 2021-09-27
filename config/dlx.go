package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func Create(cmd *cobra.Command) error {
	//make config directory and file
	err := os.MkdirAll(filepath.Join(GetConfigPath()), 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(GetConfigPath(), "dlx.yaml"))
	if err != nil {
		return err
	}
	defer f.Close()

	user, err := cmd.Flags().GetString("user")
	if err != nil {
		return errors.Wrap(err, "Error reading user")
	}

	bi, err := cmd.Flags().GetString("baseimage")
	if err != nil {
		return errors.Wrap(err, "Error reading baseimage")
	}

	sshkey, err := cmd.Flags().GetString("sshkey")
	if err != nil {
		return errors.Wrap(err, "Error reading sshkey")
	}
	if err != nil {
		return errors.Wrap(err, "Error reading lxc config. Run `lxc remote add` before configuring dlx")
	}
	lxcconf, err := GetLXCConfig()
	if err != nil {
		return errors.Wrap(err, "Error reading lxc config. Run `lxc remote add` before configuring dlx")
	}
	config := &Config{
		ClientKey:  filepath.Dir(lxcconf.Path) + "/client.key",
		ClientCert: filepath.Dir(lxcconf.Path) + "/client.crt",
		Remotes:    make(map[string]Remote),
	}

	for name, r := range lxcconf.Config.Remotes {
		if r.Protocol == "lxd" {
			config.Remotes[name] = Remote{
				User:          user,
				Host:          name,
				SSHPrivateKey: sshkey,
				BaseImage:     bi,
			}
		}
	}
	bb, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	_, err = f.Write(bb)
	if err != nil {
		return err
	}
	return nil
}

func Check() error {
	_, err := os.Stat(filepath.Join(GetConfigPath(), "dlx.yaml"))
	if err != nil {
		return err
	}
	return nil
}

func Get() (*Config, error) {
	var cfg Config
	err := Check()
	if err != nil {
		return &cfg, errors.Wrap(err, "Run 'dlx config -c' to create a default configuration file")
	}

	bb, err := ioutil.ReadFile(filepath.Join(GetConfigPath(), "dlx.yaml"))
	if err != nil {
		return &cfg, err
	}
	err = yaml.Unmarshal(bb, &cfg)
	return &cfg, err

}
