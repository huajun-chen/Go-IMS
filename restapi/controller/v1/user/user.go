package user

import (
	"Go-IMS/parameter"
	"Go-IMS/parameter/reqstruct"
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
	loginForm := reqstruct.LoginForm{}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := user.SerLogin(loginForm, c)
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

// ConUpdateUser 控制层：根据ID修改用户信息
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConUpdateUser(c *gin.Context) {
	userId := parameter.IdForm{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parError := utils.HandleValidatorError(err)
		response.Response(c, parError)
		return
	}
	// 需要修改的字段参数
	updateUserForm := reqstruct.UpdateUserForm{}
	if err := c.ShouldBindJSON(&updateUserForm); err != nil {
		parErrStr := utils.HandleValidatorError(err)
		response.Response(c, parErrStr)
		return
	}
	resStruct := user.SerUpdateUser(userId, updateUserForm, c)
	response.Response(c, resStruct)
}

// ConUpdateUserPwd 控制层：根据ID修改用户密码
// 参数：
//		c *gin.Context：gin.Context的指针
// 返回值：
//		无
func ConUpdateUserPwd(c *gin.Context) {
	userId := parameter.IdForm{}
	if err := c.ShouldBindUri(&userId); err != nil {
		parError := utils.HandleValidatorError(err)
		response.Response(c, parError)
		return
	}
	// 修改密码的参数
	updateUserPwdForm := reqstruct.UpdateUserPwdForm{}
	if err := c.ShouldBindJSON(&updateUserPwdForm); err != nil {
		parErrSrt := utils.HandleValidatorError(err)
		response.Response(c, parErrSrt)
		return
	}
	resStruct := user.SerUpdateUserPwd(userId, updateUserPwdForm, c)
	response.Response(c, resStruct)
}
