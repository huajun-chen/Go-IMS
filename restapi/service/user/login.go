package user

import (
	"Go-IMS/dao/user"
	"Go-IMS/global"
	"Go-IMS/param"
	"Go-IMS/param/req"
	"Go-IMS/param/resp"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// SerLogin 业务层：用户登录
// 参数：
//		reqLogin：登录时需要的参数
//		c：gin.Context的指针
// 返回值：
//		param.Resp：响应的结构体
func SerLogin(reqLogin req.ReqLogin, c *gin.Context) param.Resp {
	// 查询用户是否存在
	userModel, err := user.DaoGetUserByUserName(reqLogin.UserName)
	if err != nil {
		failStruct := param.Resp{
			Code: 10013,
			Msg:  global.I18nMap["10013"],
		}
		return failStruct
	}

	// 判断密码是否正确
	pwdBool := utils.CheckPassword(userModel.Password, reqLogin.Password)
	if !pwdBool {
		failStruct := param.Resp{
			Code: 10016,
			Msg:  global.I18nMap["10016"],
		}
		return failStruct
	}

	// 生成新的Token
	token := utils.CreateToken(c, userModel.ID, userModel.Role, userModel.UserName)
	data := resp.RespLogin{
		ID:    userModel.ID,
		Name:  userModel.UserName,
		Token: token,
	}
	succStruct := param.Resp{
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
//		param.Resp：响应的结构体
func SerLogout(c *gin.Context) param.Resp {
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

	succStruct := param.Resp{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2001"],
	}
	return succStruct
}
