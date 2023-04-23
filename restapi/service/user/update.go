package user

import (
	"Go-IMS/dao/user"
	"Go-IMS/global"
	"Go-IMS/model"
	"Go-IMS/parameter"
	"Go-IMS/parameter/reqstruct"
	"Go-IMS/response"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SerUpdateUser 业务层：修改用户信息
// 参数：
//		userId：用户ID
//		updateUserForm：需要修改的信息
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerUpdateUser(userId parameter.IdForm, updateInfo reqstruct.UpdateUserForm, c *gin.Context) response.ResStruct {
	// 判断是否是本人
	tokenUserId, _ := c.Get("userId")
	if userId.ID != tokenUserId {
		failStruct := response.ResStruct{
			Code: 10015,
			Msg:  global.I18nMap["10015"],
		}
		return failStruct
	}

	updateUser := model.User{
		Gender: updateInfo.Gender,
		Desc:   updateInfo.Desc,
		Mobile: updateInfo.Mobile,
		Email:  updateInfo.Email,
	}
	if err := user.DaoUpdateUser(userId.ID, updateUser); err != nil {
		failStruct := response.ResStruct{
			Code: 10003,
			Msg:  global.I18nMap["10003"],
		}
		return failStruct
	}
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2004"],
	}
	return succStruct
}

// SerUpdateUserPwd 业务层：修改用户密码
// 参数：
//		userId：用户ID
//		updateUserPwdForm：修改的密码
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerUpdateUserPwd(userId parameter.IdForm, updatePwd reqstruct.UpdateUserPwdForm, c *gin.Context) response.ResStruct {
	// 判断是否是本人
	tokenUserId, _ := c.Get("userId")
	if userId.ID != tokenUserId {
		failStruct := response.ResStruct{
			Code: 10015,
			Msg:  global.I18nMap["10015"],
		}
		return failStruct
	}
	// 查询
	userModel, err := user.DaoGetUserById(userId.ID)
	if err != nil {
		failStruct := response.ResStruct{
			Code: 10006,
			Msg:  global.I18nMap["10006"],
		}
		return failStruct
	}

	// 判断旧密码是否正确
	pwdBool := utils.CheckPassword(userModel.Password, updatePwd.PasswordOld)
	if !pwdBool {
		failStruct := response.ResStruct{
			Code: 10019,
			Msg:  global.I18nMap["10019"],
		}
		return failStruct
	}

	// 判断旧密码与新密码是否一致
	if updatePwd.PasswordOld == updatePwd.Password {
		failStruct := response.ResStruct{
			Code: 10020,
			Msg:  global.I18nMap["10020"],
		}
		return failStruct
	}

	// 判断新密码是否一致
	if updatePwd.Password != updatePwd.Password2 {
		failStruct := response.ResStruct{
			Code: 10017,
			Msg:  global.I18nMap["10017"],
		}
		return failStruct
	}

	// 密码加密
	pwdStr, _ := utils.SetPassword(updatePwd.Password)
	updateUserPwd := model.User{Password: pwdStr}
	if err := user.DaoUpdateUser(userId.ID, updateUserPwd); err != nil {
		failStruct := response.ResStruct{
			Code: 10003,
			Msg:  global.I18nMap["10003"],
		}
		return failStruct
	}
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2004"],
	}
	return succStruct
}
