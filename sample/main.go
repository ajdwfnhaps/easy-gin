package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	easygin "github.com/ajdwfnhaps/easy-gin"
	"github.com/ajdwfnhaps/easy-gin/sample/routers/api"
)

func main() {

	//创建应用程序 使用配置文件
	r := easygin.New("conf/config.toml")
	//使用swagger
	r.UseSwagger(SetSwaggerInfo)

	//注册路由
	r.RegisterRouter(api.RouterHanlder)

	//启动
	r.Run()

	//处理退出信号
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
