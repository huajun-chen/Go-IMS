package resp

// RespLogin 登录接口返回的数据
type RespLogin struct {
	ID    uint   `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
	Token string `json:"token"` // 用户Token
}

// RespUserInfo 用户信息
type RespUserInfo struct {
	ID        uint   `json:"id"`         // 用户ID
	CreatedAt string `json:"created_at"` // 创建时间
	UserName  string `json:"user_name"`  // 用户名
	Gender    string `json:"gender"`     // 性别
	Desc      string `json:"desc"`       // 描述
	Role      string `json:"role"`       // 角色
	Mobile    string `json:"mobile"`     // 电话
	Email     string `json:"email"`      // 邮箱
}

// RespUserList 用户信息列表
type RespUserList struct {
	Total  int            `json:"total"`  // 总数
	Values []RespUserInfo `json:"values"` // 用户列表
}
