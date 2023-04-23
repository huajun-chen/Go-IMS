package user

import (
	"Go-IMS/param"
	"Go-IMS/param/req"
	"Go-IMS/response"
	"Go-IMS/service/user"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
)

// ConUpdateUser 控制层：根据ID修改用户信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConUpdateUser(c *gin.Context) {
	userId := param.ReqId{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parError := utils.HandleValidatorError(err)
		response.Response(c, parError)
		return
	}
	// 需要修改的字段参数
	updateInfo := req.ReqUpdateUser{}
	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := user.SerUpdateUser(userId, updateInfo, c)
	response.Response(c, resStruct)
}

// ConUpdateUserPwd 控制层：根据ID修改用户密码
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConUpdateUserPwd(c *gin.Context) {
	userId := param.ReqId{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parError := utils.HandleValidatorError(err)
		response.Response(c, parError)
		return
	}
	// 修改密码的参数
	updatePwd := req.ReqUpdateUserPwd{}
	if err := c.ShouldBindJSON(&updatePwd); err != nil {
		parErrSrt := utils.HandleValidatorError(err)
		response.Response(c, parErrSrt)
		return
	}
	resStruct := user.SerUpdateUserPwd(userId, updatePwd, c)
	response.Response(c, resStruct)
}
