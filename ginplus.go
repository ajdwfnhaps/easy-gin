package easygin

import (
	"fmt"

	"github.com/ajdwfnhaps/easy-gin/conf"
	mw "github.com/ajdwfnhaps/easy-gin/middleware"
	"github.com/gin-gonic/gin"
)

type (
	//ConfigFunc 配置函数
	ConfigFunc func(c *conf.Config)
)

var (
	//DefaultOpts 全局配置
	DefaultOpts conf.Config
)

func init() {
	DefaultOpts = conf.Config{
		RunMode: "debug",
		HTTP: conf.HTTP{
			Host:            "0.0.0.0",
			Port:            8074,
			ShutdownTimeout: 30,
		},
		Log: conf.Log{
			LogHTTPResponse: true,
			LogHTTPRequest:  false,
		},
	}
}

// UseEasyGin 使用easy-gin插件
func UseEasyGin(optFunc ConfigFunc) error {

	//执行传入的自定义设置全局配置函数(在此之前已执行包的init方法，初始化全局配置)
	if optFunc != nil {
		optFunc(&DefaultOpts)
	}

	r := gin.New()

	r.Use(mw.LoggerMiddleware(), gin.Recovery())

	if DefaultOpts.RunMode == "debug" {
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	addr := fmt.Sprintf("%s:%d", DefaultOpts.HTTP.Host, DefaultOpts.HTTP.Port)

	go r.Run(addr)

	return nil
}
