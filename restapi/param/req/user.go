package req

// ReqLogin 用户登录参数
type ReqLogin struct {
	UserName string `json:"user_name" binding:"required,min=3,max=16"` // 用户名
	Password string `json:"password" binding:"required,min=8,max=64"`  // 密码
}

// ReqCreateUser 创建用户参数，其他信息都默认，可以让用户自己修改
type ReqCreateUser struct {
	UserName string `uri:"user_name" binding:"required,min=3,max=16"` // 用户名
}

// ReqUpdateUser 修改用户信息的参数
type ReqUpdateUser struct {
	Gender int    `json:"gender" binding:"omitempty,oneof=1 2 3"` // 性别
	Desc   string `json:"desc" binding:"omitempty,max=256"`       // 描述
	Mobile string `json:"mobile" binding:"omitempty,len=11"`      // 电话
	Email  string `json:"email" binding:"omitempty,email"`        // 邮箱
}

// ReqUpdateUserPwd 修改用户密码的参数
type ReqUpdateUserPwd struct {
	PasswordOld string `json:"password_old" binding:"required,min=8,max=64"` // 旧密码
	Password    string `json:"password" binding:"required,min=8,max=64"`     // 新密码
	Password2   string `json:"password2" binding:"required,min=8,max=64"`    // 新密码2
}
