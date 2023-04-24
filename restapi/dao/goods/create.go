package goods

import (
	"Go-WMS/global"
	"Go-WMS/model"
)

// DaoCreateGoodsCategory 创建商品分类
// 参数：
//		createInfo：需要添加的信息
// 返回值：
//		error：错误信息
func DaoCreateGoodsCategory(createInfo model.GoodsCategory) error {
	err := global.DB.Create(&createInfo).Error
	return err
}
