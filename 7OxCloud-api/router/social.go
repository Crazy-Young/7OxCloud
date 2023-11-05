package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/api"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/middleware"
)

func InitSocialRouter(Router *gin.RouterGroup) {
	SocialRouter := Router.Group("/social", middleware.JWTAuth())
	{
		SocialRouter.POST("/follow", api.Follow)
		SocialRouter.DELETE("/remove_fan", api.RemoveFan)
		SocialRouter.GET("/following", api.GetFollowing)
		SocialRouter.GET("/fan", api.GetFan)
		SocialRouter.GET("/follow_feed", api.FollowFeed)
		SocialRouter.GET("/friend_feed", api.FriendFeed)
		SocialRouter.GET("/search_following", api.SearchFollowing)
		SocialRouter.GET("/search_fan", api.SearchFan)
	}
}
