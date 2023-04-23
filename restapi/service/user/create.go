package user

import (
	"Go-WMS/dao/user"
	"Go-WMS/global"
	"Go-WMS/param"
	"Go-WMS/param/req"
	"Go-WMS/utils"
	"net/http"
)

// SerCreateUser 业务层：创建用户
// 参数：
//		reqUser：创建用户时的参数的结构体
//		c：gin.Context的指针
// 返回值：
//		param.Resp：响应的结构体
func SerCreateUser(reqUser req.ReqCreateUser) param.Resp {
	// 查询用户是否存在
	userModel, err := user.DaoGetUserByUserName(reqUser.UserName)
	if userModel != nil {
		failStruct := param.Resp{
			Code: 10014,
			Msg:  global.I18nMap["10014"],
		}
		return failStruct
	} else if err != nil {
		failStruct := param.Resp{
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
		failStruct := param.Resp{
			Code: 10006,
			Msg:  global.I18nMap["10006"],
		}
		return failStruct
	}

	// 创建用户
	err = user.DaoCreateUser(reqUser.UserName, pwdStr)
	if err != nil {
		failStruct := param.Resp{
			Code: 10018,
			Msg:  global.I18nMap["10018"],
		}
		return failStruct
	}
	succStruct := param.Resp{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2002"],
	}
	return succStruct
}
