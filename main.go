package main

import (
	"fmt"
	"os"

	"github.com/google/logger"
	"myproj.com/clmgr-lrm/config"
	"myproj.com/clmgr-lrm/pkg/common"
	"myproj.com/clmgr-lrm/pkg/agent"
	"time"
)

func main() {
	config.InitConfig()
	if err := common.InitLogger(); err != nil {
		fmt.Println("can't init logger")
		os.Exit(-1)
	}
	logger.Infof("Logger have been initialised")

	ag, err := agent.Create("test_agent1")
	if err != nil {
		logger.Errorf("Agent creation failed with error %s", err.Error())
	}
	fmt.Println(ag)

	ag.Start()
	time.Sleep(10 * time.Second)
	ag.Stop()
}