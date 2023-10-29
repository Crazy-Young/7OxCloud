package form

type LoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required_if=Type 1 ,mobile"`
	Password  string `form:"password" json:"password" binding:"required_if=Type 1,omitempty,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,len=6"`
	CaptchaId string `form:"captchaId" json:"captchaId" binding:"required_if=Type 1"`
	Type      int    `form:"type" json:"type" binding:"required,oneof=1 2"` // 1表示密码登录, 2表示验证码登录
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=20"`
	Captcha  string `form:"captcha" json:"captcha" binding:"required,len=6"`
}

type UpdateUserForm struct {
	Username        string `form:"username" json:"username" binding:"required,min=1,max=20"`
	Age             int    `form:"age" json:"age" binding:"required,gt=0"`
	Location        string `form:"location" json:"location" binding:"required,max=20"`
	Avatar          string `form:"avatar" json:"avatar" binding:"required"`
	BackgroundImage string `form:"backgroundImage" json:"backgroundImage" binding:"required"`
	Signature       string `form:"signature" json:"signature" binding:"required"`
}

type PicUrlForm struct {
	PicUrl string `form:"picUrl" json:"picUrl" binding:"required"`
}
