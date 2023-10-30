package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/api"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/middleware"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("/base")
	{
		BaseRouter.GET("/captcha", api.GetCaptcha)
		BaseRouter.POST("/send_sms", api.SendSms)

		BaseRouter.POST("upload_file", middleware.JWTAuth(), api.UploadFile)
		BaseRouter.DELETE("delete_file", middleware.JWTAuth(), api.DeleteFile)
	}
}
