package resp

type RespGoodsCategoryParent struct {
	ParentID   uint   `json:"parent_id"`   // 父级分类ID
	ParentName string `json:"parent_name"` // 父级分类名称
}

// RespGoodsCategory 商品分类信息
type RespGoodsCategory struct {
	ID           uint                    `json:"id"`            // 商品分类ID
	CreatedAt    string                  `json:"created_at"`    // 创建时间
	CategoryName string                  `json:"category_name"` // 分类名称
	ParentInfo   RespGoodsCategoryParent `json:"parent_info"`   // 父级分类信息
}

// RespGoodsCategoryList 商品分类信息列表
type RespGoodsCategoryList struct {
	Total  int                 `json:"total"`  // 总数
	Values []RespGoodsCategory `json:"values"` // 商品列表
}
