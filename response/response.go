package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	ctx.JSON(httpStatus, &Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// OK 正常响应
func OK(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, &Response{
		Code: 1,
		Data: data,
		Msg:  msg,
	})
}

// Fail 错误响应
func Fail(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, &Response{
		Code: -1,
		Data: data,
		Msg:  msg,
	})
}

// Page 分布结果响应
func Page(ctx *gin.Context, data *PagingList, msg string) {
	ctx.JSON(http.StatusOK, &Response{
		Code: 1,
		Data: &data,
		Msg:  msg,
	})
}

// Err400Response 解析请求参数发生错误响应
func Err400Response(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, &Response{
		Code: http.StatusBadRequest,
		Msg:  msg,
	})
}
