package user

import (
	"Go-IMS/param/req"
	"Go-IMS/response"
	"Go-IMS/service/user"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
)

// ConLogin 控制层：用户登录
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConLogin(c *gin.Context) {
	// 获取登录时需要的参数
	reqLogin := req.ReqLogin{}
	if err := c.ShouldBindJSON(&reqLogin); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := user.SerLogin(reqLogin, c)
	response.Response(c, resStruct)
}

// ConLogout 控制层：用户登出
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConLogout(c *gin.Context) {
	resStruct := user.SerLogout(c)
	response.Response(c, resStruct)
}
