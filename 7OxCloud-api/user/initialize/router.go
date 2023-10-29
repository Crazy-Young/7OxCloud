package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/api"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/middleware"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/router"
)

func Routers() *gin.Engine {
	if !global.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		api.HandleHttpResponse(c, http.StatusOK, "ok", nil, nil)
		return
	})

	Router.Use(middleware.Cors())

	ApiGroup := Router.Group("/api")
	router.InitUserRouter(ApiGroup)

	return Router
}
