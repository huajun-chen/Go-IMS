package goods

import (
	"Go-WMS/param/req"
	"Go-WMS/service/goods"
	"Go-WMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConCreateGoodsCategory 控制层：创建商品分类
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConCreateGoodsCategory(c *gin.Context) {
	// 获取创建商品分类需要的参数
	req := req.ReqCreateGoodsCategory{}
	if err := c.ShouldBindJSON(&req); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		c.JSON(http.StatusOK, parErrStr)
		return
	}
	resStruct := goods.SerCreateGoodsCategory(req)
	c.JSON(http.StatusOK, resStruct)
}
