package easygin

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/ajdwfnhaps/easy-gin/conf"
	"github.com/ajdwfnhaps/easy-logrus/logger"
)

var (
	configPath string
)

func TestEasyGin(t *testing.T) {

	configPath = "conf/config.toml"
	//初始化配置
	conf.Init(configPath)

	// 初始化日志
	initLog()

	cfg := conf.Global()
	log := logger.CreateLogger()
	log.Infof("iot-api服务启动，版本号：%s，进程号：%d", cfg.Version, os.Getpid())

	if err := UseEasyGin(func(opt *conf.Config) {
		opt.HTTP.Port = 8888
		opt.HTTP.ShutdownTimeout = 60
	}); err != nil {
		t.Error(err)
	}

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

func TestStartCac(t *testing.T) {
	t1 := time.Now() // get current time
	//logic handlers
	for i := 0; i < 1000; i++ {
		fmt.Print("*")
	}
	elapsed := time.Since(t1)
	a := elapsed.Milliseconds()
	fmt.Println(a)
	fmt.Println("App elapsed: ", elapsed)
}
