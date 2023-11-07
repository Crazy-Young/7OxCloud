package util

import "github.com/mojocn/base64Captcha"

var Store = base64Captcha.DefaultMemStore

func GeneratePicCaptcha() (id, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, Store)
	return cp.Generate()
}
