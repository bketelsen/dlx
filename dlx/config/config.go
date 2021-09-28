package config

type Config struct {
	Remotes map[string]Remote `yaml:"remotes"`
}

type Remote struct {
	User          string `yaml:"user"`
	Host          string `yaml:"host"`
	SSHPrivateKey string `yaml:"ssh_private_key"`
	BaseImage     string `yaml:"baseimage"`
}
