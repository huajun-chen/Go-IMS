package user

import (
	"Go-IMS/parameter"
	"Go-IMS/response"
	"Go-IMS/service/user"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
)

// ConDeleteUser 控制层：删除用户（管理员权限）
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConDeleteUser(c *gin.Context) {
	userId := parameter.IdForm{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parError := utils.HandleValidatorError(err)
		response.Response(c, parError)
		return
	}
	resStruct := user.SerDeleteUser(userId)
	response.Response(c, resStruct)
}
