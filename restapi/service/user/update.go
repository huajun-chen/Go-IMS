package user

import (
	"Go-IMS/dao/user"
	"Go-IMS/global"
	"Go-IMS/model"
	"Go-IMS/param"
	"Go-IMS/param/req"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SerUpdateUser 业务层：修改用户信息
// 参数：
//		userId：用户ID
//		reqUpdateUser：需要修改的信息
//		c：gin.Context的指针
// 返回值：
//		param.Resp：响应的结构体
func SerUpdateUser(userId param.ReqId, reqUpdateUser req.ReqUpdateUser, c *gin.Context) param.Resp {
	// 判断是否是本人
	tokenUserId, _ := c.Get("userId")
	if userId.ID != tokenUserId {
		failStruct := param.Resp{
			Code: 10015,
			Msg:  global.I18nMap["10015"],
		}
		return failStruct
	}

	updateUser := model.User{
		Gender: reqUpdateUser.Gender,
		Desc:   reqUpdateUser.Desc,
		Mobile: reqUpdateUser.Mobile,
		Email:  reqUpdateUser.Email,
	}
	if err := user.DaoUpdateUser(userId.ID, updateUser); err != nil {
		failStruct := param.Resp{
			Code: 10003,
			Msg:  global.I18nMap["10003"],
		}
		return failStruct
	}
	succStruct := param.Resp{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2004"],
	}
	return succStruct
}

// SerUpdateUserPwd 业务层：修改用户密码
// 参数：
//		userId：用户ID
//		reqUpdatePwd：修改的密码
//		c：gin.Context的指针
// 返回值：
//		param.Resp：响应的结构体
func SerUpdateUserPwd(userId param.ReqId, reqUpdatePwd req.ReqUpdateUserPwd, c *gin.Context) param.Resp {
	// 判断是否是本人
	tokenUserId, _ := c.Get("userId")
	if userId.ID != tokenUserId {
		failStruct := param.Resp{
			Code: 10015,
			Msg:  global.I18nMap["10015"],
		}
		return failStruct
	}
	// 查询
	userModel, err := user.DaoGetUserById(userId.ID)
	if err != nil {
		failStruct := param.Resp{
			Code: 10006,
			Msg:  global.I18nMap["10006"],
		}
		return failStruct
	}

	// 判断旧密码是否正确
	pwdBool := utils.CheckPassword(userModel.Password, reqUpdatePwd.PasswordOld)
	if !pwdBool {
		failStruct := param.Resp{
			Code: 10019,
			Msg:  global.I18nMap["10019"],
		}
		return failStruct
	}

	// 判断旧密码与新密码是否一致
	if reqUpdatePwd.PasswordOld == reqUpdatePwd.Password {
		failStruct := param.Resp{
			Code: 10020,
			Msg:  global.I18nMap["10020"],
		}
		return failStruct
	}

	// 判断新密码是否一致
	if reqUpdatePwd.Password != reqUpdatePwd.Password2 {
		failStruct := param.Resp{
			Code: 10017,
			Msg:  global.I18nMap["10017"],
		}
		return failStruct
	}

	// 密码加密
	pwdStr, _ := utils.SetPassword(reqUpdatePwd.Password)
	updateUserPwd := model.User{Password: pwdStr}
	if err := user.DaoUpdateUser(userId.ID, updateUserPwd); err != nil {
		failStruct := param.Resp{
			Code: 10003,
			Msg:  global.I18nMap["10003"],
		}
		return failStruct
	}
	succStruct := param.Resp{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2004"],
	}
	return succStruct
}
