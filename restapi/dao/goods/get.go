package goods

import (
	"Go-WMS/global"
	"Go-WMS/model"
	"Go-WMS/utils"
)

// DaoGetGoodsCategoryList 获取商品分类列表
// 参数：
//		getType：查询的类型（all/parent）
//		page：页数
//		pageSize：每页数量
// 返回值：
//		int：总数量
//		[]model.GoodsCategory：商品分类信息列表
// 		error：错误信息
func DaoGetGoodsCategoryList(getType string, page, pageSize int) (int, []model.GoodsCategory, error) {
	var categoryCount int64
	var category []model.GoodsCategory
	// 定义查询的条件
	var whereClause string
	var args []interface{}

	switch getType {
	case "all":
		whereClause = "1=1"
	case "parent":
		whereClause = "parent_id = ?"
		args = []interface{}{0}
	case "son":
		whereClause = "parent_id > ?"
		args = []interface{}{0}
	}
	// 查询分类信息总数量
	global.DB.Find(&category).Where(whereClause, args...).Count(&categoryCount)

	offset := utils.OffsetResult(page, pageSize)
	limit := utils.LimitResult(pageSize)
	// 根据条件获取商品分类数据
	// Preload("Parent") 关联查询
	err := global.DB.Preload("Parent").Offset(offset).Limit(limit).Order("id desc").Select(
		"id",
		"created_at",
		"category_name",
		"parent_id",
	).Where(whereClause, args...).Find(&category).Error
	if err != nil {
		return 0, nil, err
	}
	return int(categoryCount), category, nil
}
