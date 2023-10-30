package initialize

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/global"
	myvalidator "github.com/palp1tate/7OxCloud/7OxCloud-api/validator"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Translator, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0}为非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
}
