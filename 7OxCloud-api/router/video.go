package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/api"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/middleware"
)

func InitVideoRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("/video")
	{
		VideoRouter.GET("/feed", middleware.JWTAuth(), api.Feed)
		VideoRouter.POST("/publish", middleware.JWTAuth(), api.Publish)
		VideoRouter.GET("/get_video", middleware.JWTAuth(), api.GetVideo)
		VideoRouter.GET("/publish_list", middleware.JWTAuth(), api.PublishList)
		VideoRouter.GET("/topic_list", middleware.JWTAuth(), api.TopicList)
		VideoRouter.GET("/category_list", middleware.JWTAuth(), api.CategoryList)
		VideoRouter.GET("/feed_by_topic", middleware.JWTAuth(), api.FeedByTopic)
		VideoRouter.GET("/feed_by_category", middleware.JWTAuth(), api.FeedByCategory)
		VideoRouter.GET("/feed_by_search", middleware.JWTAuth(), api.FeedBySearch)
	}
}
