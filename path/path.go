package path

import (
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

const configDirName = "devlx"

func GetHomePath() string {
	// Find home directory form env
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("Get Home Dir: " + err.Error())
	}

	return home
}

func GetConfigPath() string {

	//Find the default config directory
	configPath := os.Getenv("XDG_CONFIG_HOME")
	if len(configPath) == 0 {
		configPath = filepath.Join(GetHomePath(), ".config")
	}

	//set the devlx config directory
	devlxConfigPath := filepath.Join(configPath, configDirName)

	return devlxConfigPath
}
