package common

import (
	"github.com/op/go-logging"
	"io"
	"os"
	"fmt"
)

var Logger = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func InitLogger() error {
	f, err := os.Create("/var/log/cluster-manager.log")
	if err != nil {
		fmt.Println("Can't create log file")
		return err
	}
	writer := io.Writer(f)
	backend := logging.NewLogBackend(writer, "",0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLevel := logging.AddModuleLevel(backend)
	backendLevel.SetLevel(logging.ERROR, "")

	logging.SetBackend(backendLevel, backendFormatter)

	return nil
}
