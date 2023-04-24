package goods

import (
	"Go-WMS/dao/goods"
	"Go-WMS/global"
	"Go-WMS/param"
	"Go-WMS/param/resp"
	"Go-WMS/utils"
	"net/http"
)

// SerGetGoodsCategoryList 业务层：获取商品分类列表
// 参数：
//		getType：查询的类型（all/parent）
//		reqPage：默认的页数，每页数量参数
// 返回值：
//		param.Resp：响应的结构体
func SerGetGoodsCategoryList(getType string, reqPage param.ReqPage) param.Resp {
	page, pageSize := utils.PageZero(reqPage.Page, reqPage.PageSize)
	total, goodsCategoryList, err := goods.DaoGetGoodsCategoryList(getType, page, pageSize)
	if err != nil {
		failStruct := param.Resp{
			Code: 10004,
			Msg:  global.I18nMap["10004"],
		}
		return failStruct
	}

	// 获取数据为空
	if total == 0 {
		failStruct := param.Resp{
			Code: 10005,
			Msg:  global.I18nMap["10005"],
		}
		return failStruct
	}

	// 过滤商品分类信息
	var values []resp.RespGoodsCategory
	for _, g := range goodsCategoryList {
		// 关联查询有数据，此数据有父级分类ID
		if g.Parent != nil {
			goodsCategory := resp.RespGoodsCategory{
				ID:           g.ID,
				CreatedAt:    g.CreatedAt.Format("2006-01-02"),
				CategoryName: g.CategoryName,
				ParentInfo: resp.RespGoodsCategoryParent{
					ParentID:   uint(g.ParentID), // 数据库类型int -> uint
					ParentName: g.Parent.CategoryName,
				},
			}
			values = append(values, goodsCategory)
		} else {
			// 关联查询无数据，此分类没有父级分类ID
			goodsCategory := resp.RespGoodsCategory{
				ID:           g.ID,
				CreatedAt:    g.CreatedAt.Format("2006-01-02"),
				CategoryName: g.CategoryName,
				ParentInfo: resp.RespGoodsCategoryParent{
					ParentID:   0,
					ParentName: "",
				},
			}
			values = append(values, goodsCategory)
		}
	}

	data := resp.RespGoodsCategoryList{
		Total:  total,
		Values: values,
	}
	succStruct := param.Resp{
		Code: http.StatusOK,
		Data: data,
	}
	return succStruct
}
