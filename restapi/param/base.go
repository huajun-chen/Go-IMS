package param

// ReqPage 页数，每页数量
type ReqPage struct {
	Page     int `form:"page" binding:"omitempty,gte=1,lte=10000"`      // 页数，第几页
	PageSize int `form:"page_size" binding:"omitempty,gte=1,lte=10000"` // 每页的数量
}

// ReqId ID
type ReqId struct {
	ID uint `uri:"id" binding:"required,gte=1,lte=100000000"` // 主键ID
}

// Resp 响应结构体
type Resp struct {
	// omitempty(省略)：字段如果没有值就不显示此字段
	Code int         `json:"code,omitempty"` // 自定义响应状态码
	Msg  string      `json:"msg,omitempty"`  // 响应信息
	Data interface{} `json:"data,omitempty"` // 响应数据
}
