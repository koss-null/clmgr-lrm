package config

import (
	"github.com/BurntSushi/toml"
	"myproj.com/clmgr-lrm/pkg/common"
	"os"
)

const configFile = "/opt/clmgr/config/config.toml"

type config struct {
	AgentPath string
}

var Config config

func InitConfig() {
	if _, err := toml.DecodeFile(configFile, &Config); err != nil {
		common.Logger.Error("Can't parse config, error %s", err.Error())
		os.Exit(1)
	}
}
