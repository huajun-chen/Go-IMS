package resstruct

// LoginReturn 登录接口返回的数据
type LoginReturn struct {
	ID    uint   `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
	Token string `json:"token"` // 用户Token
}

// UserInfoReturn 用户信息
type UserInfoReturn struct {
	ID        uint   `json:"id"`         // 用户ID
	CreatedAt string `json:"created_at"` // 创建时间
	UserName  string `json:"user_name"`  // 用户名
	Gender    string `json:"gender"`     // 性别
	Desc      string `json:"desc"`       // 描述
	Role      string `json:"role"`       // 角色
	Mobile    string `json:"mobile"`     // 电话
	Email     string `json:"email"`      // 邮箱
}

// UserInfoListReturn 用户信息列表
type UserInfoListReturn struct {
	Total  int              `json:"total"`  // 总数
	Values []UserInfoReturn `json:"values"` // 用户列表
}
