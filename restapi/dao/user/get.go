package user

import (
	"Go-IMS/global"
	"Go-IMS/model"
	"Go-IMS/utils"
	"gorm.io/gorm"
)

// DaoGetUserById 根据用户ID查询用户是否存在，并返回用户信息
// 参数：
//		userId：用户ID
// 返回值：
//		*model.User：用户信息的指针
//		error：错误信息
func DaoGetUserById(userId uint) (*model.User, error) {
	var user model.User
	res := global.DB.First(&user, userId)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			// 用户不存在
			return nil, nil
		} else {
			// 查询出错
			return nil, res.Error
		}
	}
	// 返回用户信息
	return &user, nil
}

// DaoGetUserByUserName 根据用户名查询用户是否存在，并返回用户信息
// 参数：
//		userName：用户名
// 返回值：
//		*model.User：用户信息的指针
//		error：错误信息
func DaoGetUserByUserName(userName string) (*model.User, error) {
	var user model.User
	res := global.DB.Where("user_name = ?", userName).First(&user)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			// 用户不存在
			return nil, nil
		} else {
			// 查询出错
			return nil, res.Error
		}
	}
	// 返回用户信息
	return &user, nil
}

// DaoGetUserList 获取用户列表信息
// 参数：
//		page：页数
//		pageSize：每页数量
// 返回值：
//		int：总数量
//		[]user.User：用户信息列表
// 		error：错误信息
func DaoGetUserList(page, pageSize int) (int, []model.User, error) {
	var usersCount int64
	var users []model.User
	// 查询用户总数量
	global.DB.Find(&users).Count(&usersCount)

	offset := utils.OffsetResult(page, pageSize)
	limit := utils.LimitResult(pageSize)
	// 根据条件获取用户数据
	err := global.DB.Offset(offset).Limit(limit).Order("id desc").Select(
		"id",
		"created_at",
		"user_name",
		"gender",
		"desc",
		"role",
		"mobile",
		"email").Find(&users).Error
	if err != nil {
		return 0, nil, err
	}
	return int(usersCount), users, nil
}
