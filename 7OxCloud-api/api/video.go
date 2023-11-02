package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/form"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/middleware"
	"github.com/palp1tate/7OxCloud/proto/video"
	"github.com/palp1tate/7OxCloud/util"
)

func Feed(c *gin.Context) {
	latestTime := util.String2Int64(c.Query("latestTime"))
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	currentUserId := c.GetInt("userId")
	res, err := global.VideoServiceClient.Feed(c, &videoProto.FeedRequest{
		LatestTime:    latestTime,
		CurrentUserId: int64(currentUserId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	videos := make([]gin.H, len(res.Videos))
	for i, video := range res.Videos {
		videos[i] = MiniVideoToResponse(video)
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取视频流成功", refreshedToken, gin.H{
		"videos":   videos,
		"nextTime": res.NextTime,
	})
}

func HotFeed(c *gin.Context) {
	latestTime := util.String2Int64(c.Query("latestTime"))
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	currentUserId := c.GetInt("userId")
	res, err := global.VideoServiceClient.HotFeed(c, &videoProto.FeedRequest{
		LatestTime:    latestTime,
		CurrentUserId: int64(currentUserId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	videos := make([]gin.H, len(res.Videos))
	for i, video := range res.Videos {
		videos[i] = MiniVideoToResponse(video)
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取热门视频流成功", refreshedToken, gin.H{
		"videos":   videos,
		"nextTime": res.NextTime,
	})
}

func Publish(c *gin.Context) {
	publishForm := form.PublishForm{}
	if err := c.ShouldBind(&publishForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	_, err := global.VideoServiceClient.Publish(c, &videoProto.PublishRequest{
		Description: publishForm.Description,
		PlayUrl:     publishForm.PlayUrl,
		CoverUrl:    publishForm.CoverUrl,
		UserId:      int64(c.GetInt("userId")),
		Topics:      publishForm.Topics,
		CategoryId:  int64(publishForm.CategoryId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "发布视频成功", refreshedToken, nil)
}

func GetVideo(c *gin.Context) {
	videoId := util.String2Int64(c.Query("vid"))
	if videoId == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "vid不能为空", nil, nil)
		return
	}
	currentUserId := c.GetInt("userId")
	res, err := global.VideoServiceClient.Video(c, &videoProto.VideoRequest{
		Id:            videoId,
		CurrentUserId: int64(currentUserId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取视频成功", refreshedToken, VideoToResponse(res.Video))
}

func PublishList(c *gin.Context) {
	latestTime := util.String2Int64(c.Query("latestTime"))
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	uid := util.String2Int64(c.Query("uid"))
	currentUserId := c.GetInt("userId")
	if uid == 0 {
		uid = int64(currentUserId)
	}
	res, err := global.VideoServiceClient.PublishList(c, &videoProto.PublishListRequest{
		UserId:     uid,
		LatestTime: latestTime,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	videos := make([]gin.H, len(res.Videos))
	for i, video := range res.Videos {
		videos[i] = SmallVideoToResponse(video)
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取视频发布列表成功", refreshedToken, gin.H{
		"videos":   videos,
		"nextTime": res.NextTime,
	})
}

func TopicList(c *gin.Context) {
	res, err := global.VideoServiceClient.TopicList(c, &empty.Empty{})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	currentUserId := c.GetInt("userId")
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取话题列表成功", refreshedToken, res.Topics)
}

func CategoryList(c *gin.Context) {
	res, err := global.VideoServiceClient.CategoryList(c, &empty.Empty{})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	currentUserId := c.GetInt("userId")
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取分类列表成功", refreshedToken, res.Categories)
}

func FeedByTopic(c *gin.Context) {
	topicId := util.String2Int64(c.Query("tid"))
	if topicId == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "tid不能为空", nil, nil)
		return
	}
	latestTime := util.String2Int64(c.Query("latestTime"))
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	currentUserId := c.GetInt("userId")
	res, err := global.VideoServiceClient.FeedByTopic(c, &videoProto.FeedByTopicRequest{
		LatestTime:    latestTime,
		TopicId:       topicId,
		CurrentUserId: int64(currentUserId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	videos := make([]gin.H, len(res.Videos))
	for i, video := range res.Videos {
		videos[i] = VideoToResponse(video)
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取视频流成功", refreshedToken, gin.H{
		"videos":   videos,
		"nextTime": res.NextTime,
	})
}

func FeedByCategory(c *gin.Context) {
	categoryId := util.String2Int64(c.Query("categoryId"))
	if categoryId == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "categoryId不能为空", nil, nil)
		return
	}
	latestTime := util.String2Int64(c.Query("latestTime"))
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	currentUserId := c.GetInt("userId")
	res, err := global.VideoServiceClient.FeedByCategory(c, &videoProto.FeedByCategoryRequest{
		LatestTime:    latestTime,
		CategoryId:    categoryId,
		CurrentUserId: int64(currentUserId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	videos := make([]gin.H, len(res.Videos))
	for i, video := range res.Videos {
		videos[i] = MiniVideoToResponse(video)
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取视频流成功", refreshedToken, gin.H{
		"videos":   videos,
		"nextTime": res.NextTime,
	})
}

func FeedBySearch(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		HandleHttpResponse(c, http.StatusBadRequest, "keyword不能为空", nil, nil)
		return
	}
	latestTime := util.String2Int64(c.Query("latestTime"))
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	currentUserId := c.GetInt("userId")
	res, err := global.VideoServiceClient.FeedBySearch(c, &videoProto.FeedBySearchRequest{
		LatestTime:    latestTime,
		Keyword:       keyword,
		CurrentUserId: int64(currentUserId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	videos := make([]gin.H, len(res.Videos))
	for i, video := range res.Videos {
		videos[i] = VideoToResponse(video)
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取视频流成功", refreshedToken, gin.H{
		"videos":   videos,
		"nextTime": res.NextTime,
	})
}

func DeleteVideo(c *gin.Context) {
	videoId := util.String2Int64(c.Query("vid"))
	if videoId == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "vid不能为空", nil, nil)
		return
	}
	currentUserId := c.GetInt("userId")
	_, err := global.VideoServiceClient.DeleteVideo(c, &videoProto.VideoRequest{
		Id:            videoId,
		CurrentUserId: int64(currentUserId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "删除视频成功", refreshedToken, nil)
}

func User2Response(user *videoProto.User) gin.H {
	return gin.H{
		"uid":      user.Id,
		"username": user.Username,
		"avatar":   user.Avatar,
		"location": user.Location,
		"isFollow": user.IsFollow,
	}
}

func MiniUserToResponse(user *videoProto.MiniUser) gin.H {
	return gin.H{
		"uid":      user.Id,
		"username": user.Username,
		"isFollow": user.IsFollow,
	}
}

func VideoToResponse(video *videoProto.Video) gin.H {
	return gin.H{
		"vid":          strconv.FormatInt(video.Id, 10),
		"description":  video.Description,
		"playUrl":      video.PlayUrl,
		"coverUrl":     video.CoverUrl,
		"likeCount":    video.LikeCount,
		"commentCount": video.CommentCount,
		"collectCount": video.CollectCount,
		"isLike":       video.IsLike,
		"isCollect":    video.IsCollect,
		"author":       User2Response(video.Author),
		"topics":       video.Topics,
		"publishTime":  video.PublishTime,
	}
}

func MiniVideoToResponse(video *videoProto.MiniVideo) gin.H {
	return gin.H{
		"vid":         strconv.FormatInt(video.Id, 10),
		"playUrl":     video.PlayUrl,
		"coverUrl":    video.CoverUrl,
		"likeCount":   video.LikeCount,
		"description": video.Description,
		"topics":      video.Topics,
		"author":      MiniUserToResponse(video.Author),
		"publishTime": video.PublishTime,
	}
}

func SmallVideoToResponse(video *videoProto.SmallVideo) gin.H {
	return gin.H{
		"vid":         strconv.FormatInt(video.Id, 10),
		"playUrl":     video.PlayUrl,
		"coverUrl":    video.CoverUrl,
		"likeCount":   video.LikeCount,
		"description": video.Description,
		"topics":      video.Topics,
	}
}
