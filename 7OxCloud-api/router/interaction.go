package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/api"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/middleware"
)

func InitInteractionRouter(Router *gin.RouterGroup) {
	InteractionRouter := Router.Group("/interaction", middleware.JWTAuth())
	{
		InteractionRouter.POST("/like_video", api.LikeVideo)
		InteractionRouter.POST("/collect_video", api.CollectVideo)
		InteractionRouter.GET("/like_list", api.GetLikeList)
		InteractionRouter.GET("/collect_list", api.GetCollectList)
		InteractionRouter.POST("/comment", api.Comment)
		InteractionRouter.DELETE("/delete_comment", api.DeleteComment)
		InteractionRouter.GET("/comment_list", api.GetCommentList)
	}
}
