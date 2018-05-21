package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"fmt"
)

const configFile = "/opt/clmgr/config/config.toml"

// const configFile = "config/config-debug.toml"

// 4 testing inside pkg
// const configFile = "../../config/config-debug-pkg.toml"

type config struct {
	AgentPath          string
	ResourcesPath      string
	LogPath            string
	CoordinatorAddress string
}

var Config config

func InitConfig() {
	if _, err := toml.DecodeFile(configFile, &Config); err != nil {
		fmt.Printf("Can't parse config, error %s\n", err.Error())
		os.Exit(1)
	}
}
