package user

import (
	"Go-IMS/global"
	"Go-IMS/model"
)

// DaoUpdateUser 根据用户ID更新用户信息
// 参数：
//		userId：用户ID
//		updateUser：需要更新的信息
// 返回值：
//		error：错误信息
func DaoUpdateUser(userId uint, updateUser model.User) error {
	err := global.DB.Model(&model.User{}).Where("id = ?", userId).Updates(updateUser).Error
	return err
}
