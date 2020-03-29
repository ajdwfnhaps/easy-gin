package api

import (
	"github.com/gin-gonic/gin"
)

//RouterHanlder 定义所有api路由
func RouterHanlder(app *gin.Engine) error {

	g := app.Group("/api")

	v1 := g.Group("/v1")
	{
		// 注册/api/v1/demos
		gDemo := v1.Group("demos")
		{
			cDemo := NewDemoController()
			// gDemo.GET("", cDemo.Query)
			gDemo.GET(":id", cDemo.Get)
			gDemo.POST("", cDemo.Create)
			// gDemo.PUT(":id", cDemo.Update)
			// gDemo.DELETE(":id", cDemo.Delete)
			// gDemo.PATCH(":id/enable", cDemo.Enable)
			// gDemo.PATCH(":id/disable", cDemo.Disable)

			// // @description This is the first line
			// // @description This is the second line
			// // @description And so forth.
			// gDemo.GET("/ping", func(c *gin.Context) {
			// 	log := logger.CreateLogger()
			// 	log.WithField("myKey", "myValue").Printf("log test...")
			// 	c.JSON(200, gin.H{
			// 		"message": "pong",
			// 	})
			// })
		}
	}

	return nil
}
