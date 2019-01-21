package config

import (
	yaml "gopkg.in/yaml.v2"
)

type CloudInit struct {
	RunCommands []Command `yaml:"runcmd"`
	Packages    []string  `yaml:"packages"`
}

func (ci *CloudInit) Append(other CloudInit) {
	ci.Packages = append(ci.Packages, other.Packages...)
	ci.RunCommands = append(ci.RunCommands, other.RunCommands...)
}

func Read(bb []byte) (CloudInit, error) {
	var ci CloudInit
	err := yaml.Unmarshal(bb, &ci)
	if err != nil {
		return ci, err
	}
	return ci, nil
}

type Command []string

func (a *Command) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var multi []string
	err := unmarshal(&multi)
	if err != nil {
		var single string
		err := unmarshal(&single)
		if err != nil {
			return err
		}
		*a = []string{single}
	} else {
		*a = multi
	}
	return nil
}
