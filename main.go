package main

import (
	"fmt"
	"os"

	"myproj.com/clmgr-lrm/config"
	"myproj.com/clmgr-lrm/pkg/common"
)

func main() {
	err := common.InitLogger()
	if err != nil {
		fmt.Println("can't init logger")
		os.Exit(-1)
	}
	config.InitConfig()
}