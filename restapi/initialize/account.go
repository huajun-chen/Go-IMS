package initialize

import (
	"Go-WMS/dao/user"
	"Go-WMS/global"
	"Go-WMS/model"
	"Go-WMS/utils"
)

// InitAdminAccount 初始化一个admin账户
// 参数：
//		无
// 返回值：
//		无
func InitAdminAccount() {
	// 默认配置的管理员用户名
	adminInfo := global.Settings.AdminInfo
	// 查询admin是否存在
	userModel, _, err := user.DaoGetUserByUserName(adminInfo.UserName)
	// 不存在，创建
	if userModel == nil && err == nil {
		// 加密密码
		pwdStr, err := utils.SetPassword(adminInfo.Password)
		if err != nil {
			panic(err)
		}
		// 创建管理员账户
		user := model.User{UserName: adminInfo.UserName, Password: pwdStr, Role: 1}
		if err := global.DB.Create(&user).Error; err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
}
