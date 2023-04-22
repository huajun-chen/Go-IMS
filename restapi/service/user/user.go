package user

import (
	dao "Go-IMS/dao/user"
	"Go-IMS/global"
	"Go-IMS/model/user"
	"Go-IMS/parameter"
	"Go-IMS/parameter/reqstruct"
	"Go-IMS/parameter/resstruct"
	"Go-IMS/response"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

// SerLogin 业务层：用户登录
// 参数：
//		loginForm：登录时需要的参数
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerLogin(loginForm reqstruct.LoginForm, c *gin.Context) response.ResStruct {
	// 查询用户是否存在
	userInfo, ok := dao.DaoFindUserInfoToUserName(loginForm.UserName)
	if !ok {
		failStruct := response.ResStruct{
			Code: 10013,
			Msg:  global.I18nMap["10013"],
		}
		return failStruct
	}
	// 判断密码是否正确
	pwdBool := utils.CheckPassword(userInfo.Password, loginForm.Password)
	if !pwdBool {
		failStruct := response.ResStruct{
			Code: 10016,
			Msg:  global.I18nMap["10016"],
		}
		return failStruct
	}
	// 生成新的Token
	token := utils.CreateToken(c, userInfo.ID, userInfo.Role, userInfo.UserName)
	data := resstruct.LoginReturn{
		ID:    userInfo.ID,
		Name:  userInfo.UserName,
		Token: token,
	}
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2000"],
		Data: data,
	}
	return succStruct
}

// SerLogout 业务层：用户登出
// 参数：
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerLogout(c *gin.Context) response.ResStruct {
	// 获取Token
	tokenStr, _ := c.Get("token")
	// 获取用户ID
	userId, _ := c.Get("userId")
	// 获取Token到期时间
	tokenExpiresAt, _ := c.Get("tokenExpiresAt")
	// 计算Token剩余的时间（Token到期时间戳 - 当前时间戳）
	timeLeft := time.Duration(tokenExpiresAt.(int64)-time.Now().Unix()) * time.Second
	// 计算Token MD5值
	tokenMD5 := utils.MD5(tokenStr.(string))
	// 将Key（Token MD5值），value（用户ID），到期时间（Token剩余的时间）加入Redis
	// 延迟10秒执行，避免此用户的其他请求还未返回Token就失效
	go func() {
		time.Sleep(10 * time.Second)
		err := utils.RedisSetStr(tokenMD5, userId, timeLeft)
		if err != nil {
			// Set Redis 错误的话，只记录日志
			zap.S().Errorf("token MD5 Set Redis faild：%s", tokenStr)
		}
	}()

	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2001"],
	}
	return succStruct
}

// SerCreateUser 业务层：创建用户
// 参数：
//		createUserForm：创建用户时的参数的结构体
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerCreateUser(createUserForm reqstruct.CreateUserForm) response.ResStruct {
	// 查询用户是否存在
	_, ok := dao.DaoFindUserInfoToUserName(createUserForm.UserName)
	if ok {
		failStruct := response.ResStruct{
			Code: 10014,
			Msg:  global.I18nMap["10014"],
		}
		return failStruct
	}
	// 创建用户
	createUserOK := dao.CreateUserInfo(createUserForm.UserName)
	if !createUserOK {
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

// SerDeleteUser 业务层：根据用户ID删除用户信息
// 参数：
//		userId：用户ID
// 返回值：
//		response.ResStruct：响应的结构体
func SerDeleteUser(userId parameter.IdForm) response.ResStruct {
	// 删除用户
	err := dao.DaoDeleteUserInfoToId(userId.ID)
	if err != nil {
		failStruct := response.ResStruct{
			Code: 10002,
			Msg:  global.I18nMap["10002"],
		}
		return failStruct
	}
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2003"],
	}
	return succStruct
}

// SerGetUserList 业务层：获取用户列表
// 参数：
//		pageForm：默认的页数，每页数量参数
// 返回值：
//		response.ResStruct：响应的结构体
func SerGetUserList(pageForm parameter.PageForm) response.ResStruct {
	page, pageSize := utils.PageZero(pageForm.Page, pageForm.PageSize)
	total, userList, err := dao.DaoGetUserList(page, pageSize)
	if err != nil {
		failStruct := response.ResStruct{
			Code: 10004,
			Msg:  global.I18nMap["10004"],
		}
		return failStruct
	}
	// 获取数据为空
	if total == 0 {
		failStruct := response.ResStruct{
			Code: 10005,
			Msg:  global.I18nMap["10005"],
		}
		return failStruct
	}
	// 过滤用户信息
	var values []resstruct.UserInfoReturn
	for _, u := range userList {
		userInfo := resstruct.UserInfoReturn{
			ID:        u.ID,
			CreatedAt: u.CreatedAt.Format("2006-01-02"),
			UserName:  u.UserName,
			Gender:    strconv.Itoa(u.Gender),
			Desc:      u.Desc,
			Role:      strconv.Itoa(u.Role),
			Mobile:    u.Mobile,
			Email:     u.Email,
		}
		values = append(values, userInfo)
	}
	data := resstruct.UserInfoListReturn{
		Total:  total,
		Values: values,
	}
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Data: data,
	}
	return succStruct
}

// SerGetUser 业务层：根据ID获取用户信息
// 参数：
//		userId：用户ID
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerGetUser(userId parameter.IdForm, c *gin.Context) response.ResStruct {
	// 判断是否是本人
	tokenUserId, _ := c.Get("userId")
	if tokenUserId != userId.ID {
		failStruct := response.ResStruct{
			Code: 10015,
			Msg:  global.I18nMap["10015"],
		}
		return failStruct
	}
	// 通过用户ID获取信息，查询的ID和Token的ID一致，说明ID一定能正常登录，数据肯定存在
	myselfInfo, _ := dao.DaoFindUserInfoToId(userId.ID)
	data := resstruct.UserInfoReturn{
		ID:        myselfInfo.ID,
		CreatedAt: myselfInfo.CreatedAt.Format("2006-01-02"),
		UserName:  myselfInfo.UserName,
		Gender:    strconv.Itoa(myselfInfo.Gender),
		Desc:      myselfInfo.Desc,
		Role:      strconv.Itoa(myselfInfo.Role),
		Mobile:    myselfInfo.Mobile,
		Email:     myselfInfo.Email,
	}
	succStruct := response.ResStruct{
		Code: http.StatusOK,
		Data: data,
	}
	return succStruct
}

// SerUpdateUser 业务层：修改用户信息
// 参数：
//		userId：用户ID
//		updateUserForm：需要修改的信息
//		c：gin.Context的指针
// 返回值：
//		response.ResStruct：响应的结构体
func SerUpdateUser(userId parameter.IdForm, updateUserForm reqstruct.UpdateUserForm, c *gin.Context) response.ResStruct {
	// 判断是否是本人
	tokenUserId, _ := c.Get("userId")
	if userId.ID != tokenUserId {
		failStruct := response.ResStruct{
			Code: 10015,
			Msg:  global.I18nMap["10015"],
		}
		return failStruct
	}

	updateUser := user.User{
		Gender: updateUserForm.Gender,
		Desc:   updateUserForm.Desc,
		Mobile: updateUserForm.Mobile,
		Email:  updateUserForm.Email,
	}
	if err := dao.DaoUpdateUserInfo(userId.ID, updateUser); err != nil {
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
func SerUpdateUserPwd(userId parameter.IdForm, updateUserPwdForm reqstruct.UpdateUserPwdForm, c *gin.Context) response.ResStruct {
	// 判断是否是本人
	tokenUserId, _ := c.Get("userId")
	if userId.ID != tokenUserId {
		failStruct := response.ResStruct{
			Code: 10015,
			Msg:  global.I18nMap["10015"],
		}
		return failStruct
	}
	// 判断旧密码是否正确
	userInfo, _ := dao.DaoFindUserInfoToId(userId.ID)
	pwdBool := utils.CheckPassword(userInfo.Password, updateUserPwdForm.PasswordOld)
	if !pwdBool {
		failStruct := response.ResStruct{
			Code: 10019,
			Msg:  global.I18nMap["10019"],
		}
		return failStruct
	}
	// 判断旧密码与新密码是否一致
	if updateUserPwdForm.PasswordOld == updateUserPwdForm.Password {
		failStruct := response.ResStruct{
			Code: 10020,
			Msg:  global.I18nMap["10020"],
		}
		return failStruct
	}
	// 判断新密码是否一致
	if updateUserPwdForm.Password != updateUserPwdForm.Password2 {
		failStruct := response.ResStruct{
			Code: 10017,
			Msg:  global.I18nMap["10017"],
		}
		return failStruct
	}
	// 密码加密
	pwdStr, _ := utils.SetPassword(updateUserPwdForm.Password)
	updateUserPwd := user.User{Password: pwdStr}
	if err := dao.DaoUpdateUserInfo(userId.ID, updateUserPwd); err != nil {
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
