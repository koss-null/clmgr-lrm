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
	if err := common.InitLogger(); err != nil {
		fmt.Println("can't init logger")
		os.Exit(-1)
	}
	logger.Infof("Logger have been initialised")

	ag, err := configuration.CreateAgent("test_agent1")
	if err != nil {
		logger.Errorf("Agent creation failed with error %s", err.Error())
	}
	fmt.Println(ag)
}