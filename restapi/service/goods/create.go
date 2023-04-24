package goods

import (
	"Go-WMS/dao/goods"
	"Go-WMS/global"
	"Go-WMS/model"
	"Go-WMS/param"
	"Go-WMS/param/req"
	"net/http"
)

// SerCreateGoodsCategory 业务层：创建商品分类
// 参数：
//		req：创建商品分类的参数
// 返回值：
//		param.Resp：响应的结构体
func SerCreateGoodsCategory(req req.ReqCreateGoodsCategory) param.Resp {
	createInfo := model.GoodsCategory{
		CategoryName: req.CategoryName,
		ParentID:     req.ParentID,
	}
	if err := goods.DaoCreateGoodsCategory(createInfo); err != nil {
		failStruct := param.Resp{
			Code: 10001,
			Msg:  global.I18nMap["10001"],
		}
		return failStruct
	}

	succStruct := param.Resp{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2002"],
	}
	return succStruct
}
