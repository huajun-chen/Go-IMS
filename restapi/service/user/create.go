package user

import (
	"Go-IMS/dao/user"
	"Go-IMS/global"
	"Go-IMS/parameter/reqstruct"
	"Go-IMS/response"
	"Go-IMS/utils"
	"net/http"
)

// SerCreateUser 业务层：创建用户
// 参数：
//		createUserForm：创建用户时的参数的结构体
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerCreateUser(userForm reqstruct.CreateUserForm) response.ResStruct {
	// 查询用户是否存在
	userModel, err := user.DaoGetUserByUserName(userForm.UserName)
	if userModel != nil {
		failStruct := response.ResStruct{
			Code: 10014,
			Msg:  global.I18nMap["10014"],
		}
		return failStruct
	} else if err != nil {
		failStruct := response.ResStruct{
			Code: 10006,
			Msg:  global.I18nMap["10006"],
		}
		return failStruct
	}

	// 密码使用admin账户的默认密码
	password := global.Settings.AdminInfo.Password
	// 密码加密
	pwdStr, err := utils.SetPassword(password)
	if err != nil {
		failStruct := response.ResStruct{
			Code: 10006,
			Msg:  global.I18nMap["10006"],
		}
		return failStruct
	}

	// 创建用户
	err = user.DaoCreateUser(userForm.UserName, pwdStr)
	if err != nil {
		failStruct := response.ResStruct{
			Code: 10018,
			Msg:  global.I18nMap["10018"],
		}
		return failStruct
	}
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2002"],
	}
	return succStruct
}
