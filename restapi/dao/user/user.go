package dao

import (
	"Go-IMS/global"
	"Go-IMS/model/user"
	"Go-IMS/utils"
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

// DaoFindUserInfoToId 根据用ID查询用户是否存在，并返回用户信息
// 参数：
//		userId：用户ID
// 返回值：
//		*model.User：用户信息的指针
//		bool：查询是否成功
func DaoFindUserInfoToId(userId uint) (*user.User, bool) {
	var userInfo user.User
	rows := global.DB.First(&userInfo, userId)
	if rows.RowsAffected < 1 {
		return &userInfo, false
	}
	return &userInfo, true
}

// CreateUserInfo 管理员权限创建用户
// 参数：
//		userName：用户名
// 返回值：
//		bool：创建是否成功
func CreateUserInfo(userName string) bool {
	// 密码使用admin账户的默认密码
	password := global.Settings.AdminInfo.Password
	// 密码加密
	pwdStr, err := utils.SetPassword(password)
	if err != nil {
		return false
	}
	userInfo := user.User{
		UserName: userName,
		Password: pwdStr,
		Role:     2, // 普通用户
	}
	if err := global.DB.Create(&userInfo).Error; err != nil {
		return false
	}
	return true
}

// DaoDeleteUserInfoToId 根据ID删除用户信息
// 参数：
//		userId：用户ID
// 返回值：
//		error：错误信息
func DaoDeleteUserInfoToId(userId uint) error {
	err := global.DB.Delete(&user.User{}, userId).Error
	return err
}

// DaoGetUserList 获取用户列表信息
// 参数：
//		page：页数
//		pageSize：每页数量
// 返回值：
//		int：总数量
//		[]user.User：用户信息列表
// 		error：错误信息
func DaoGetUserList(page, pageSize int) (int, []user.User, error) {
	var usersCount int64
	var users []user.User
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

// DaoUpdateUserInfo 根据ID更新用户信息
// 参数：
//		userId：用户ID
//		updateUser：需要更新的信息
// 返回值：
//		error：错误信息
func DaoUpdateUserInfo(userId uint, updateUser user.User) error {
	err := global.DB.Model(&user.User{}).Where("id = ?", userId).Updates(updateUser).Error
	if err != nil {
		return err
	}
	return nil
}
