package user

import (
	"Go-WMS/global"
	"Go-WMS/model"
)

// DaoCreateUser 管理员权限创建用户
// 参数：
//		userName：用户名
// 返回值：
//		bool：创建是否成功
func DaoCreateUser(userName, pwdStr string) error {
	userInfo := model.User{
		UserName: userName,
		Password: pwdStr,
		Role:     2, // 普通用户
	}
	err := global.DB.Create(&userInfo).Error
	return err
}
