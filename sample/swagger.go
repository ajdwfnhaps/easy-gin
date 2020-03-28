package main

import (
	"github.com/ajdwfnhaps/easy-gin/sample/docs"

	"github.com/ajdwfnhaps/easy-gin/conf"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// UseSwagger 使用Swagger文档
func UseSwagger(app *gin.Engine) {
	cfg := conf.Global()
	swag := cfg.Swagger

	if swag.On != 1 {
		return
	}

	// programatically set swagger info
	docs.SwaggerInfo.Title = swag.Title
	docs.SwaggerInfo.Description = swag.Description
	docs.SwaggerInfo.Version = swag.Version
	docs.SwaggerInfo.Host = swag.Host
	docs.SwaggerInfo.BasePath = swag.BasePath
	docs.SwaggerInfo.Schemes = swag.Schemes
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
