package middleware

import (
	"bytes"
	"io/ioutil"
	"mime"
	"net/http"
	"time"

	"github.com/ajdwfnhaps/easy-gin/conf"
	"github.com/ajdwfnhaps/easy-gin/response"
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
		logConf := conf.Global().Log

		if logConf.LogHTTPRequest {
			method := c.Request.Method
			// 如果是POST/PUT请求，并且内容类型为JSON，则读取内容体
			if method == http.MethodPost || method == http.MethodPut {
				mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
				if mediaType == "application/json" {
					body, err := ioutil.ReadAll(c.Request.Body)
					c.Request.Body.Close()
					if err == nil {
						buf := bytes.NewBuffer(body)
						c.Request.Body = ioutil.NopCloser(buf)
						logger = logger.WithFields(logrus.Fields{
							"content_length": c.Request.ContentLength,
							"body":           string(body),
						})
					}
				}
			}
		}

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

		if logConf.LogHTTPResponse {
			logger = logger.WithField("res_length", c.Writer.Size())

			if v, ok := c.Get(response.ResBodyKey); ok {
				if b, ok := v.([]byte); ok {
					logger = logger.WithField("res_body", string(b))
				}
			}
		}

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
