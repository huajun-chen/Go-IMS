package user

import (
	"Go-IMS/dao/user"
	"Go-IMS/global"
	"Go-IMS/param"
	"Go-IMS/param/resp"
	"Go-IMS/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SerGetUserList 业务层：获取用户列表
// 参数：
//		reqPage：默认的页数，每页数量参数
// 返回值：
//		param.Resp：响应的结构体
func SerGetUserList(reqPage param.ReqPage) param.Resp {
	page, pageSize := utils.PageZero(reqPage.Page, reqPage.PageSize)
	total, userList, err := user.DaoGetUserList(page, pageSize)
	if err != nil {
		failStruct := param.Resp{
			Code: 10004,
			Msg:  global.I18nMap["10004"],
		}
		return failStruct
	}

	// 获取数据为空
	if total == 0 {
		failStruct := param.Resp{
			Code: 10005,
			Msg:  global.I18nMap["10005"],
		}
		return failStruct
	}

	// 过滤用户信息
	var values []resp.RespUserInfo
	for _, u := range userList {
		userInfo := resp.RespUserInfo{
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
	data := resp.RespUserList{
		Total:  total,
		Values: values,
	}
	succStruct := param.Resp{
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
//		param.Resp：响应的结构体
func SerGetUser(userId param.ReqId, c *gin.Context) param.Resp {
	// 判断是否是本人
	tokenUserId, _ := c.Get("userId")
	if tokenUserId != userId.ID {
		failStruct := param.Resp{
			Code: 10015,
			Msg:  global.I18nMap["10015"],
		}
		return failStruct
	}

	// 用户能正常登录，说明用户信息一定存在
	userModel, err := user.DaoGetUserById(userId.ID)
	if err != nil {
		failStruct := param.Resp{
			Code: 10006,
			Msg:  global.I18nMap["10006"],
		}
		return failStruct
	}

	data := resp.RespUserInfo{
		ID:        userModel.ID,
		CreatedAt: userModel.CreatedAt.Format("2006-01-02"),
		UserName:  userModel.UserName,
		Gender:    strconv.Itoa(userModel.Gender),
		Desc:      userModel.Desc,
		Role:      strconv.Itoa(userModel.Role),
		Mobile:    userModel.Mobile,
		Email:     userModel.Email,
	}
	succStruct := param.Resp{
		Code: http.StatusOK,
		Data: data,
	}
	return succStruct
}
