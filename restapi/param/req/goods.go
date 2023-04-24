package req

// ReqCreateGoodsCategory 创建商品分类的参数
type ReqCreateGoodsCategory struct {
	CategoryName string `json:"category_name" binding:"required,min=1,max=16"`     // 商品分类
	ParentID     int    `json:"parent_id" binding:"omitempty,gte=1,lte=100000000"` // 父级分类ID
}

// ReqGetGoodsCategoryList 获取商品分类列表的参数
type ReqGetGoodsCategoryList struct {
	Type string `uri:"type" binding:"required,oneof='all' 'parent'"` // 查询的类型，查询全部分类/只查询父级分类
}
