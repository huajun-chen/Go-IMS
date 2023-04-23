package user

import (
	"Go-IMS/parameter"
	"Go-IMS/response"
	"Go-IMS/service/user"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
)

// ConGetUser 控制层：根据ID查询用户信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetUser(c *gin.Context) {
	userId := parameter.IdForm{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parError := utils.HandleValidatorError(err)
		response.Response(c, parError)
		return
	}
	resStruct := user.SerGetUser(userId, c)
	response.Response(c, resStruct)
}

// ConGetUserList 控制层：用户列表（管理员权限）
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConGetUserList(c *gin.Context) {
	// 页数，页码参数
	pageForm := parameter.PageForm{}
	if err := c.ShouldBindQuery(&pageForm); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := user.SerGetUserList(pageForm)
	response.Response(c, resStruct)
}
