package reqstruct

// LoginForm 用户登录参数
type LoginForm struct {
	UserName string `json:"user_name" binding:"required,min=3,max=16"` // 用户名
	Password string `json:"password" binding:"required,min=8,max=64"`  // 密码
}
