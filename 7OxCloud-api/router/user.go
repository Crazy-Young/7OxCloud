package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/api"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{
		UserRouter.POST("/register", api.Register)
		UserRouter.POST("/login", api.Login)
		UserRouter.GET("/get_user", middleware.JWTAuth(), api.GetUser)
		UserRouter.PUT("/update_user", middleware.JWTAuth(), api.UpdateUser)
		UserRouter.PUT("/reset_password", api.ResetPassword)
	}
}
