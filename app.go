package easygin

import (
	"fmt"
	"os"

	"github.com/ajdwfnhaps/easy-gin/conf"
	mw "github.com/ajdwfnhaps/easy-gin/middleware"
	"github.com/ajdwfnhaps/easy-logrus/logger"
	"github.com/gin-gonic/gin"
)

//App 应用程序
type App struct {
	ConfPath string
	Opts     *conf.Config
	Gin      *gin.Engine
	Logger   *logger.Entry
	isUseLog bool
}

//New 创建应用程序
func New(fPath string) *App {
	r := gin.New()
	r.Use(gin.Recovery())
	app := &App{Gin: r}

	app.Opts = &DefaultOpts

	if len(fPath) > 0 {
		app = app.WithConf(fPath)
	}
	return app
}

//WithConf 使用配置文件
func (c *App) WithConf(fPath string) *App {
	c.ConfPath = fPath
	conf.Init(fPath)
	c.Opts = conf.Global()
	if c.Opts.Log.AppNo != 0 {
		c = c.withLogrusConf()
	}
	return c
}

//withLogrusConf 使用logrus日志组件
func (c *App) withLogrusConf() *App {
	if err := logger.UseLogrusWithConfig(c.ConfPath); err != nil {
		panic("logger init failed, " + err.Error())
	}
	c.isUseLog = true
	c.Gin.Use(mw.LoggerMiddleware())
	return c
}

//Run 启动应用
func (c *App) Run() {
	if c.isUseLog {
		c.Logger = logger.CreateLogger()
		c.Logger.Infof("easy-gin启动，版本号：%s，进程号：%d", c.Opts.Version, os.Getpid())
	}

	if c.Opts.RunMode == "debug" {
		c.Gin.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	addr := fmt.Sprintf("%s:%d", c.Opts.HTTP.Host, c.Opts.HTTP.Port)
	go c.Gin.Run(addr)
}
