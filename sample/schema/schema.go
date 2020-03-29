package schema

// HTTPStatusText 定义HTTP状态文本
type HTTPStatusText string

func (t HTTPStatusText) String() string {
	return string(t)
}

// 定义HTTP状态文本常量
const (
	OKStatusText HTTPStatusText = "OK"
)

// HTTPError HTTP响应错误
type HTTPError struct {
	Error HTTPErrorItem `json:"error"` // 错误项
}

// HTTPErrorItem HTTP响应错误项
type HTTPErrorItem struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
}

// HTTPStatus HTTP响应状态
type HTTPStatus struct {
	Status string `json:"status"` // 状态(OK)
}

// HTTPList HTTP响应列表数据
type HTTPList struct {
	List       interface{}     `json:"list"`
	Pagination *HTTPPagination `json:"pagination,omitempty"`
}

// HTTPPagination HTTP分页数据
type HTTPPagination struct {
	Total    int `json:"total"`
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}

// PaginationParam 分页查询条件
type PaginationParam struct {
	PageIndex int // 页索引
	PageSize  int // 页大小
}

// PaginationResult 分页查询结果
type PaginationResult struct {
	Total int // 总数据条数
}

// Result 响应结果
type Result struct {
	Success bool        `json:"success"`
	Title   string      `json:"title"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// PagingList 分页列表
type PagingList struct {
	PageIndex      int         `json:"pageIndex"`
	PageSize       int         `json:"pageSize"`
	TotalCount     int         `json:"totalCount"`
	PageTotalCount int         `json:"pageTotalCount"`
	Items          interface{} `json:"items"`
}
