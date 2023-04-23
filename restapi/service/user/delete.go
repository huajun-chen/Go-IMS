package user

import (
	"Go-IMS/dao/user"
	"Go-IMS/global"
	"Go-IMS/param"
	"net/http"
)

// SerDeleteUser 业务层：根据用户ID删除用户信息
// 参数：
//		userId：用户ID
// 返回值：
//		param.Resp：响应的结构体
func SerDeleteUser(userId param.ReqId) param.Resp {
	err := user.DaoDeleteUserById(userId.ID)
	if err != nil {
		failStruct := param.Resp{
			Code: 10002,
			Msg:  global.I18nMap["10002"],
		}
		return failStruct
	}
	succStruct := param.Resp{
		Code: http.StatusOK,
		Msg:  global.I18nMap["2003"],
	}
	return succStruct
}
