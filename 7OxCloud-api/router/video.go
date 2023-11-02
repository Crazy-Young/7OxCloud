package router

import (
	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/api"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/middleware"
)

func InitVideoRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("/video", middleware.JWTAuth())
	{
		VideoRouter.GET("/feed", api.Feed)
		VideoRouter.POST("/publish", api.Publish)
		VideoRouter.GET("/get_video", api.GetVideo)
		VideoRouter.GET("/publish_list", api.PublishList)
		VideoRouter.GET("/topic_list", api.TopicList)
		VideoRouter.GET("/category_list", api.CategoryList)
		VideoRouter.GET("/feed_by_topic", api.FeedByTopic)
		VideoRouter.GET("/feed_by_category", api.FeedByCategory)
		VideoRouter.GET("/feed_by_search", api.FeedBySearch)
		VideoRouter.GET("/hot_feed", api.HotFeed)
		VideoRouter.DELETE("/delete_video", api.DeleteVideo)
	}
}
