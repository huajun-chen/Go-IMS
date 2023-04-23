package utils

import (
	"Go-IMS/global"
	"Go-IMS/model"
)

// Migration MYSQL迁移，建表，更新表
// 参数：
//		无
// 返回值：
//		无
func Migration() {
	_ = global.DB.AutoMigrate(&model.User{}) // 用户
}
