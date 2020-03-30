package easygin

import (
	"fmt"
	"os"

	"github.com/ajdwfnhaps/easy-gin/conf"
	mw "github.com/ajdwfnhaps/easy-gin/middleware"
	"github.com/ajdwfnhaps/easy-logrus/logger"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

//Default 默认配置应用
func Default(fPath string) *App {
	app := New(fPath)
	app.Gin.NoMethod(mw.NoMethodHandler())
	app.Gin.NoRoute(mw.NoRouteHandler())
	return app
}

//WithConf 使用配置文件
func (c *App) WithConf(fPath string) *App {
	c.ConfPath = fPath
	conf.Init(fPath)
	c.Opts = conf.Global()
	return c
}

//UseLogrusConf 使用logrus日志组件
func (c *App) UseLogrusConf(skippers ...mw.SkipperFunc) *App {
	if err := logger.UseLogrusWithConfig(c.ConfPath); err != nil {
		panic("logger init failed, " + err.Error())
	}
	c.isUseLog = true
	c.Gin.Use(mw.LoggerMiddleware(skippers...))
	return c
}

//Run 启动应用
func (c *App) Run() {
	if c.isUseLog {
		c.Logger = logger.CreateLogger()
		c.Logger.Infof("easy-gin启动，版本号：%s，进程号：%d", c.Opts.Version, os.Getpid())
	}

	//添加测试路由
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

//UseSwagger 使用swagger
func (c *App) UseSwagger(setSwagInfo func()) *App {
	if setSwagInfo != nil {
		setSwagInfo()
	}
	c.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return c
}

//RegisterRouter 注册路由
func (c *App) RegisterRouter(routerHanlder func(*gin.Engine) error) *App {

	if routerHanlder != nil {
		if err := routerHanlder(c.Gin); err != nil {
			panic(err)
		}
	}
	return c
}
