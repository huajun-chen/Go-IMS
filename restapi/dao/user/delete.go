package user

import (
	"Go-IMS/global"
	"Go-IMS/model"
)

// DaoDeleteUserById 根据ID删除用户信息
// 参数：
//		userId：用户ID
// 返回值：
//		error：错误信息
func DaoDeleteUserById(userId uint) error {
	err := global.DB.Delete(&model.User{}, userId).Error
	return err
}
