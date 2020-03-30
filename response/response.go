package response

import (
	"encoding/json"
	"net/http"

	"github.com/ajdwfnhaps/easy-gin/conf"
	"github.com/gin-gonic/gin"
)

const (
	prefix = "easy-gin"
	// UserIDKey 存储上下文中的键(用户ID)
	UserIDKey = prefix + "/user-id"
	// TraceIDKey 存储上下文中的键(跟踪ID)
	TraceIDKey = prefix + "/trace-id"
	// ResBodyKey 存储上下文中的键(响应Body数据)
	ResBodyKey = prefix + "/res-body"
)

// Response 响应体结构
type Response struct {
	Code int         `json:"code,omitempty"` //自定义业务code
	Data interface{} `json:"data,omitempty"` //返回数据结果
	Msg  string      `json:"msg,omitempty"`  //提示信息
}

// PagingList 分页列表
type PagingList struct {
	PageIndex      int         `json:"pageIndex"`
	PageSize       int         `json:"pageSize"`
	TotalCount     int         `json:"totalCount"`
	PageTotalCount int         `json:"pageTotalCount"`
	Items          interface{} `json:"items"`
}

// Result 结果响应
func Result(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {

	v := &Response{
		Code: code,
		Data: data,
		Msg:  msg,
	}

	cfg := conf.Global()
	if cfg.Log.LogHTTPResponse {
		buf, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		ctx.Set(ResBodyKey, buf)
		ctx.Data(httpStatus, "application/json; charset=utf-8", buf)
	} else {
		ctx.JSON(httpStatus, v)
	}
	ctx.Abort()
}

// OK 正常响应
func OK(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, http.StatusOK, 1, data, msg)
}

// Fail 错误响应
func Fail(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, http.StatusOK, -1, data, msg)
}

// Page 分布结果响应
func Page(ctx *gin.Context, data *PagingList, msg string) {
	Result(ctx, http.StatusOK, 1, &data, msg)
}

// ErrResponse 错误响应
func ErrResponse(ctx *gin.Context, httpStatus int, msg string) {
	Result(ctx, httpStatus, httpStatus, nil, msg)
}

// Err400Response 解析请求参数发生错误响应
func Err400Response(ctx *gin.Context, msg string) {
	Result(ctx, http.StatusBadRequest, http.StatusBadRequest, nil, msg)
}
