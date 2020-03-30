package middleware

import (
	"time"

	"github.com/ajdwfnhaps/easy-logrus/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//LoggerMiddleware 日志中间件
func LoggerMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		logger := logger.CreateLogger()

		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		// 执行时间
		latencyTime := time.Since(startTime)
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
			"elapsed_time": latencyTime.Milliseconds(),
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqURL,
		}).Info()

	}
}
