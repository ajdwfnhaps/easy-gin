package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	easygin "github.com/ajdwfnhaps/easy-gin"
	mw "github.com/ajdwfnhaps/easy-gin/middleware"
	"github.com/ajdwfnhaps/easy-gin/sample/routers/api"
)

func main() {

	//创建应用程序 使用配置文件
	//r := easygin.New("conf/config.toml")
	r := easygin.Default("conf/config.toml")

	//使用跨域请求中间件
	r.UseCors()

	//使用logrus日志组件
	//指定api路径规则才记录日志
	apiPrefixes := []string{"/api/"}
	r.UseLogrusConf(mw.AllowPathPrefixNoSkipper(apiPrefixes...))

	//注册路由
	r.RegisterRouter(api.RouterHanlder)

	//使用swagger
	r.UseSwagger(SetSwaggerInfo)

	//使用静态站点中间件
	r.UseWWWRoot()

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
