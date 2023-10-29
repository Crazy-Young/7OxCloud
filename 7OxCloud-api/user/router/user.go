package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/api"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{
		UserRouter.GET("/captcha", api.GetCaptcha)
		UserRouter.POST("/send_sms", api.SendSms)

		UserRouter.POST("upload_picture", middleware.JWTAuth(), api.UploadPicture)
		UserRouter.DELETE("delete_picture", middleware.JWTAuth(), api.DeletePicture)

		UserRouter.POST("/register", api.Register)
		UserRouter.POST("/login", api.Login)
		UserRouter.GET("/get_user", middleware.JWTAuth(), api.GetUser)
		UserRouter.PUT("/update_user", middleware.JWTAuth(), api.UpdateUser)
	}
}
