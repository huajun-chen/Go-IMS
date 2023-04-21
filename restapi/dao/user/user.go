package dao

import (
	"Go-IMS/global"
	"Go-IMS/model/user"
)

// DaoFindUserInfoToUserName 根据用户名查询用户是否存在，并返回用户信息
// 参数：
//		userName：用户名
// 返回值：
//		*model.User：用户信息的指针
//		bool：查询是否成功
func DaoFindUserInfoToUserName(userName string) (*user.User, bool) {
	var userInfo user.User
	// 查询用户
	rows := global.DB.Where(&user.User{UserName: userName}).Find(&userInfo)
	if rows.RowsAffected < 1 {
		return &userInfo, false
	}
	return &userInfo, true
}
