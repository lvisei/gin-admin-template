package schema

// LoginParam 登录参数
type LoginParam struct {
	UserName    string `json:"userName" binding:"required"`    // 用户名
	Password    string `json:"password" binding:"required"`    // 密码(md5加密)
	CaptchaID   string `json:"captchaId" binding:"required"`   // 验证码ID
	CaptchaCode string `json:"captchaCode" binding:"required"` // 验证码
}

// UserLoginInfo 用户登录信息
type UserLoginInfo struct {
	UserID   string `json:"userId"`   // 用户ID
	UserName string `json:"userName"` // 用户名
	RealName string `json:"realName"` // 真实姓名
	Phone    string `json:"phone"`    // 手机号
	Email    string `json:"email"`    // 邮箱
	Avatar   string `json:"avatar"`   // 头像
	Roles    Roles  `json:"roles"`    // 角色列表
}

// UserLoginInfo 用户登录信息
type UpdateUserParam struct {
	UserName string `json:"userName" binding:"required"` // 用户名
	RealName string `json:"realName" binding:"required"` // 真实姓名
	Phone    string `json:"phone"`                       // 手机号
	Email    string `json:"email"`                       // 邮箱
	Avatar   string `json:"avatar"`                      // 头像
}

// UpdatePasswordParam 更新密码请求参数
type UpdatePasswordParam struct {
	OldPassword string `json:"oldPassword" binding:"required"` // 旧密码(md5加密)
	NewPassword string `json:"newPassword" binding:"required"` // 新密码(md5加密)
}

// LoginCaptcha 登录验证码
type LoginCaptcha struct {
	CaptchaID string `json:"captchaId"` // 验证码ID
}

// LoginTokenInfo 登录令牌信息
type LoginTokenInfo struct {
	AccessToken string `json:"accessToken"` // 访问令牌
	TokenType   string `json:"tokenType"`   // 令牌类型
	ExpiresAt   int64  `json:"expiresAt"`   // 令牌到期时间戳
}
