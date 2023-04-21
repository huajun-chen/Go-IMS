package resstruct

// LoginReturn 登录接口返回的数据
type LoginReturn struct {
	ID    uint   `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
	Token string `json:"token"` // 用户Token
}
