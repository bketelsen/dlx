package config

import (
	"os"
	"path/filepath"

	lxcconfig "github.com/lxc/lxd/lxc/config"
)

type LXC struct {
	Config *lxcconfig.Config
	Path   string
}

func GetLXCConfig() (*LXC, error) {
	c := &LXC{}
	cfg, path, err := load()
	if err != nil {
		return nil, err
	}
	c.Config = cfg
	c.Path = path
	return c, nil
}
func (c *LXC) DefaultRemote() lxcconfig.Remote {
	return c.Config.Remotes[c.Config.DefaultRemote]
}

func (c *LXC) GetRemotes() map[string]lxcconfig.Remote {
	var remotes = make(map[string]lxcconfig.Remote)
	for key, r := range c.Config.Remotes {
		if r.Protocol == "lxd" {
			remotes[key] = r
		}
	}
	return remotes
}

func (c *LXC) Save() error {
	return c.Config.SaveConfig(c.Path)
}

func load() (*lxcconfig.Config, string, error) {
	paths := []string{
		"snap/lxd/common/config/config.yml",
		".config/lxd/config.yml",
	}
	found := false
	configPath := ""
	for _, path := range paths {
		cfgp := filepath.Join(GetHomePath(), path)
		if _, err := os.Stat(cfgp); err == nil {
			found = true
			configPath = cfgp
			break
		}
	}
	if !found {
		return nil, "", os.ErrNotExist
	}
	cfg, err := lxcconfig.LoadConfig(configPath)
	return cfg, configPath, err

}
