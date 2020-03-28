package sample

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	easygin "github.com/ajdwfnhaps/easy-gin"
	"github.com/ajdwfnhaps/easy-gin/conf"
)

func TestEasyGin(t *testing.T) {
	if err := easygin.UseEasyGin(func(opt *conf.Config) {
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
