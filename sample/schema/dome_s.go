package schema

import (
	"time"
)

// Demo demo对象
type Demo struct {
	RecordID  string    `json:"record_id"`                             // 记录ID
	Code      string    `json:"code" binding:"required"`               // 编号
	Name      string    `json:"name" binding:"required"`               // 名称
	Memo      string    `json:"memo"`                                  // 备注
	Status    int       `json:"status" binding:"required,max=2,min=1"` // 状态(1:启用 2:停用)
	Creator   string    `json:"creator"`                               // 创建者
	CreatedAt time.Time `json:"created_at"`                            // 创建时间
}

// DemoQueryParam 查询条件
type DemoQueryParam struct {
	Code     string `json:"code,omitempty"`      // 编号
	Status   int    `json:"status,omitempty"`    // 状态(1:启用 2:停用)
	LikeCode string `json:"like_code,omitempty"` // 编号(模糊查询)
	LikeName string `json:"like_name,omitempty"` // 名称(模糊查询)
}
