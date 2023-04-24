package goods

import (
	"Go-WMS/param"
	"Go-WMS/param/req"
	"Go-WMS/service/goods"
	"Go-WMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConGetGoodsCategoryList 控制层：商品分类列表
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetGoodsCategoryList(c *gin.Context) {
	// 获取uri中的查询类型的参数
	getType := req.ReqGetGoodsCategoryList{}
	if err := c.ShouldBindUri(&getType); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		c.JSON(http.StatusOK, parErrStr)
		return
	}
	// 页数，页码参数
	reqPage := param.ReqPage{}
	if err := c.ShouldBindQuery(&reqPage); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		c.JSON(http.StatusOK, parErrStr)
		return
	}
	resStruct := goods.SerGetGoodsCategoryList(getType.Type, reqPage)
	c.JSON(http.StatusOK, resStruct)
}
