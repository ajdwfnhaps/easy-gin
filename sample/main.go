package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	easygin "github.com/ajdwfnhaps/easy-gin"
	"github.com/ajdwfnhaps/easy-logrus/logger"
)

var (
	configPath string
)

func main() {

	r := easygin.New("conf/config.toml")
	r.Run()

	handleSignal()
}

func handleSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-c:
		fmt.Println("服务退出")
	}
}

func initLog() {
	if err := logger.UseLogrusWithConfig(configPath); err != nil {
		panic("logger init failed, " + err.Error())
	}
}
