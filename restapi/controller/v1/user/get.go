package user

import (
	"Go-WMS/param"
	"Go-WMS/service/user"
	"Go-WMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConGetUser 控制层：根据ID查询用户信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetUser(c *gin.Context) {
	userId := param.ReqId{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parError := utils.HandleValidatorError(err)
		c.JSON(http.StatusOK, parError)
		return
	}
	resStruct := user.SerGetUser(userId, c)
	c.JSON(http.StatusOK, resStruct)
}

// ConGetUserList 控制层：用户列表（管理员权限）
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetUserList(c *gin.Context) {
	// 页数，页码参数
	reqPage := param.ReqPage{}
	if err := c.ShouldBindQuery(&reqPage); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		c.JSON(http.StatusOK, parErrStr)
		return
	}
	resStruct := user.SerGetUserList(reqPage)
	c.JSON(http.StatusOK, resStruct)
}
