package config

type Config struct {
	Host          string   `yaml:"host"`
	User          string   `yaml:"user"`
	Socket        string   `yaml:"socket"`
	BaseImage     string   `yaml:"baseimage"`
	ClientCert    string   `yaml:"clientcert"`
	ClientKey     string   `yaml:"clientkey"`
	Profiles      []string `yaml:"profiles"`
	SSHPrivateKey string   `yaml:"ssh_private_key"`
}
