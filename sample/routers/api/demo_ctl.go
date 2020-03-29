package api

import (
	"time"

	"github.com/ajdwfnhaps/easy-gin/sample/schema"
	"github.com/gin-gonic/gin"
)

//DemoController demo控制器
type DemoController struct {
}

//NewDemoController 创建demo控制器
func NewDemoController() *DemoController {
	return &DemoController{}
}

/*
Tags 是用来给API分组的。
@Param 1.参数名 2.参数类型 3.参数数据类型 4.是否必须 5.参数描述 6.其他属性
*/

// Get 获取单个demo
// @Tags Demo
// @Summary 查询获取单个demo数据
// @Description 向你说描述
// @Param id path integer true "id"
// @Success 200 {object} schema.Demo
// @Router /api/v1/demos/{id} [get]
func (d *DemoController) Get(c *gin.Context) {

	id := c.Param("id")
	c.JSON(200, &schema.Demo{
		RecordID:  id,
		Code:      id,
		Name:      "成员国",
		Status:    1,
		Creator:   "admin",
		CreatedAt: time.Now(),
	})
}

// Create 创建demo
// @Tags Demo
// @Summary 创建单个demo
// @Description 向你说描述
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.Demo true "创建数据"
// @Success 200 {object} schema.Result
// @Router /api/v1/demos [post]
func (d *DemoController) Create(c *gin.Context) {
	c.JSON(200, &schema.Result{
		Success: true,
		Title:   "创建demo",
		Msg:     "创建成功",
	})
}
