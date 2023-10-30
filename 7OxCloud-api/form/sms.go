package form

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Type   int    `form:"type" json:"type" binding:"required,oneof=1 2 3"` // 1表示注册, 2表示登录, 3表示修改密码
}
