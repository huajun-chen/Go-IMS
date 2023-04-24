package req

// ReqCreateGoodsCategory 创建商品分类的参数
type ReqCreateGoodsCategory struct {
	CategoryName string `json:"category_name" binding:"required,min=1,max=16"`     // 商品分类
	ParentID     int    `json:"parent_id" binding:"omitempty,gte=1,lte=100000000"` // 父级分类ID
}
