package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/consts"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/form"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/middleware"
	"github.com/palp1tate/7OxCloud/proto/social"
	"github.com/palp1tate/7OxCloud/util"
)

func Follow(c *gin.Context) {
	followForm := form.FollowForm{}
	if err := c.ShouldBind(&followForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	fanId := c.GetInt("userId")
	if fanId == followForm.UserId {
		HandleHttpResponse(c, http.StatusBadRequest, "不能对自己进行该操作", nil, nil)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	switch followForm.Type {
	case consts.Follow:
		_, err := global.SocialServiceClient.Follow(c, &socialProto.FollowRequest{
			FanId:  int64(fanId),
			UserId: int64(followForm.UserId),
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		HandleHttpResponse(c, http.StatusOK, "关注成功", refreshedToken, nil)

	case consts.CancelFollow:
		_, err := global.SocialServiceClient.CancelFollow(c, &socialProto.CancelFollowRequest{
			FanId:  int64(fanId),
			UserId: int64(followForm.UserId),
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		HandleHttpResponse(c, http.StatusOK, "取消关注成功", refreshedToken, nil)
	}
	return
}

func RemoveFan(c *gin.Context) {
	fanId := util.String2Int64(c.Query("fanId"))
	if fanId == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "fanId不能为空", nil, nil)
		return
	}
	currentUserId := int64(c.GetInt("userId"))
	_, err := global.SocialServiceClient.CancelFollow(c, &socialProto.CancelFollowRequest{
		FanId:  fanId,
		UserId: currentUserId,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "移除粉丝成功", refreshedToken, nil)
}

func GetFollowing(c *gin.Context) {
	uid := util.String2Int64(c.Query("uid"))
	currentUserId := int64(c.GetInt("userId"))
	if uid == 0 {
		uid = currentUserId
	}
	res, err := global.SocialServiceClient.GetFollowing(c, &socialProto.GetFollowingRequest{
		CurrentUserId: uid,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	followings := make([]gin.H, len(res.Followings))
	for i, following := range res.Followings {
		followings[i] = FollowingToResponse(following)
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取关注列表成功", refreshedToken, gin.H{
		"followings": followings,
		"count":      res.Count,
	})
}

func SearchFollowing(c *gin.Context) {
	currentUserId := int64(c.GetInt("userId"))
	keyword := c.Query("keyword")
	if keyword == "" {
		HandleHttpResponse(c, http.StatusBadRequest, "keyword不能为空", nil, nil)
		return
	}
	res, err := global.SocialServiceClient.SearchFollowing(c, &socialProto.SearchRequest{
		CurrentUserId: currentUserId,
		Keyword:       keyword,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	followings := make([]gin.H, len(res.Followings))
	for i, following := range res.Followings {
		followings[i] = FollowingToResponse(following)
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "搜索关注列表成功", refreshedToken, followings)
	return
}

func GetFan(c *gin.Context) {
	uid := util.String2Int64(c.Query("uid"))
	currentUserId := int64(c.GetInt("userId"))
	if uid == 0 {
		uid = currentUserId
	}
	res, err := global.SocialServiceClient.GetFan(c, &socialProto.GetFanRequest{
		CurrentUserId: uid,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	fans := make([]gin.H, len(res.Fans))
	for i, fan := range res.Fans {
		fans[i] = FanToResponse(fan)
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取粉丝列表成功", refreshedToken, gin.H{
		"fans":  res.Fans,
		"count": res.Count,
	})
}

func SearchFan(c *gin.Context) {
	currentUserId := int64(c.GetInt("userId"))
	keyword := c.Query("keyword")
	if keyword == "" {
		HandleHttpResponse(c, http.StatusBadRequest, "keyword不能为空", nil, nil)
		return
	}
	res, err := global.SocialServiceClient.SearchFan(c, &socialProto.SearchRequest{
		CurrentUserId: currentUserId,
		Keyword:       keyword,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	fans := make([]gin.H, len(res.Fans))
	for i, fan := range res.Fans {
		fans[i] = FanToResponse(fan)
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "搜索粉丝列表成功", refreshedToken, fans)
	return
}

func FollowFeed(c *gin.Context) {
	latestTime := util.String2Int64(c.Query("latestTime"))
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	uid := c.GetInt("userId")
	res, err := global.SocialServiceClient.FollowFeed(c, &socialProto.FollowFeedRequest{
		UserId:     int64(uid),
		LatestTime: latestTime,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	videos := make([]gin.H, len(res.Videos))
	for i, video := range res.Videos {
		videos[i] = Video2Response(video)
	}

	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "获取关注动态成功", refreshedToken, gin.H{
		"videos":   videos,
		"nextTime": res.NextTime,
	})
}

func FriendFeed(c *gin.Context) {
	latestTime := util.String2Int64(c.Query("latestTime"))
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	uid := c.GetInt("userId")
	res, err := global.SocialServiceClient.FriendFeed(c, &socialProto.FriendFeedRequest{
		UserId:     int64(uid),
		LatestTime: latestTime,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	videos := make([]gin.H, len(res.Videos))
	for i, video := range res.Videos {
		videos[i] = Video2Response(video)
	}

	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "获取好友动态成功", refreshedToken, gin.H{
		"videos":   videos,
		"nextTime": res.NextTime,
	})
}

func Video2Response(video *socialProto.Video) gin.H {
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
		"topics":       video.Topics,
		"author":       video.Author,
		"publishTime":  video.PublishTime,
	}
}

func FollowingToResponse(following *socialProto.Following) gin.H {
	return gin.H{
		"uid":       following.Id,
		"username":  following.Username,
		"avatar":    following.Avatar,
		"signature": following.Signature,
		"isFan":     following.IsFan,
	}
}

func FanToResponse(fan *socialProto.Fan) gin.H {
	return gin.H{
		"uid":       fan.Id,
		"username":  fan.Username,
		"avatar":    fan.Avatar,
		"signature": fan.Signature,
		"isFollow":  fan.IsFollow,
	}
}
