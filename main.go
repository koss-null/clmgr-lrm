package main

import (
	"fmt"
	"os"

	"myproj.com/clmgr-lrm/config"
	"myproj.com/clmgr-lrm/pkg/common"
	"myproj.com/clmgr-lrm/pkg/configuration"
	"github.com/google/logger"
)

func main() {
	config.InitConfig()
	err := common.InitLogger()
	if err != nil {
		fmt.Println("can't init logger")
		os.Exit(-1)
	}
	logger.Info("Logger have been initialised")

	_, err = configuration.CreateAgent("test_agent1")
	if err != nil {
		logger.Error("Agent creation failed with error %s", err.Error())
	}
}