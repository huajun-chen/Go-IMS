package user

import (
	"Go-IMS/parameter/reqstruct"
	"Go-IMS/response"
	"Go-IMS/service/user"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
)

// ConCreateUser 控制层：创建用户（管理员权限）
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConCreateUser(c *gin.Context) {
	// 从uri中获取用户名
	userName := reqstruct.CreateUserForm{}
	if err := c.ShouldBindUri(&userName); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := user.SerCreateUser(userName)
	response.Response(c, resStruct)
}
