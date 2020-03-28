package easygin

import (
	"fmt"

	"github.com/ajdwfnhaps/easy-gin/conf"
	"github.com/gin-gonic/gin"
)

type (
	//ConfigFunc 配置函数
	ConfigFunc func(c *conf.Config)
)

var (
	//GlobalGinOption 全局配置
	GlobalGinOption conf.Config
)

func init() {
	GlobalGinOption = conf.Config{
		RunMode: "debug",
		HTTP: conf.HTTP{
			Host:            "0.0.0.0",
			Port:            8074,
			ShutdownTimeout: 30,
		},
	}
}

// UseEasyGin 使用easy-gin插件
func UseEasyGin(optFunc ConfigFunc) error {

	//执行传入的自定义设置全局配置函数(在此之前已执行包的init方法，初始化全局配置)
	if optFunc != nil {
		optFunc(&GlobalGinOption)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	addr := fmt.Sprintf("%s:%d", GlobalGinOption.HTTP.Host, GlobalGinOption.HTTP.Port)

	r.Run(addr)

	return nil
}
