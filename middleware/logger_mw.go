package middleware

import (
	"time"

	"github.com/ajdwfnhaps/easy-logrus/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {

	//-----------初始化日志组件
	// if err := logger.UseLogrus(func(c *conf.LogOption) {

	// 	c.AppName = "Go应用002"
	// 	c.LogFileRotationTime = 60
	// 	c.LogFilePathFormat = ".%Y-%m-%d.log"

	// }); err != nil {
	// 	fmt.Println(err)
	// }

	logger := logger.CreateLogger()

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqURL := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqURL,
		}).Info()

	}
}
