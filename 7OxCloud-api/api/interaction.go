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
	"github.com/palp1tate/7OxCloud/proto/interaction"
	"github.com/palp1tate/7OxCloud/util"
)

func ViewVideo(c *gin.Context) {
	vid := util.String2Int64(c.Query("vid"))
	if vid == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "vid不能为空", nil, nil)
		return
	}
	_, err := global.InteractionServiceClient.ViewVideo(c, &interactionProto.ViewVideoRequest{
		UserId:  int64(c.GetInt("userId")),
		VideoId: vid,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "观看视频成功", refreshedToken, nil)
}

func LikeVideo(c *gin.Context) {
	likeVideoForm := form.LikeVideoForm{}
	if err := c.ShouldBind(&likeVideoForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	switch likeVideoForm.Type {
	case consts.LikeVideo:
		_, err := global.InteractionServiceClient.LikeVideo(c, &interactionProto.LikeVideoRequest{
			UserId:  int64(c.GetInt("userId")),
			VideoId: util.String2Int64(likeVideoForm.VideoId),
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		HandleHttpResponse(c, http.StatusOK, "点赞成功", refreshedToken, nil)
	case consts.CancelLikeVideo:
		_, err := global.InteractionServiceClient.CancelLikeVideo(c, &interactionProto.CancelLikeVideoRequest{
			UserId:  int64(c.GetInt("userId")),
			VideoId: util.String2Int64(likeVideoForm.VideoId),
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		HandleHttpResponse(c, http.StatusOK, "取消点赞成功", refreshedToken, nil)
	}
	return
}

func CollectVideo(c *gin.Context) {
	collectVideoForm := form.CollectVideoForm{}
	if err := c.ShouldBind(&collectVideoForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	switch collectVideoForm.Type {
	case consts.CollectVideo:
		_, err := global.InteractionServiceClient.CollectVideo(c, &interactionProto.CollectVideoRequest{
			UserId:  int64(c.GetInt("userId")),
			VideoId: util.String2Int64(collectVideoForm.VideoId),
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		HandleHttpResponse(c, http.StatusOK, "收藏成功", refreshedToken, nil)
	case consts.CancelCollectVideo:
		_, err := global.InteractionServiceClient.CancelCollectVideo(c, &interactionProto.CancelCollectVideoRequest{
			UserId:  int64(c.GetInt("userId")),
			VideoId: util.String2Int64(collectVideoForm.VideoId),
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		HandleHttpResponse(c, http.StatusOK, "取消收藏成功", refreshedToken, nil)
	}
	return
}

func GetLikeList(c *gin.Context) {
	latestTime := util.String2Int64(c.Query("latestTime"))
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	uid := util.String2Int64(c.Query("uid"))
	currentUserId := c.GetInt("userId")
	if uid == 0 {
		uid = int64(currentUserId)
	}
	res, err := global.InteractionServiceClient.GetLikeList(c, &interactionProto.GetLikeListRequest{
		UserId:     uid,
		LatestTime: latestTime,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	videos := make([]gin.H, len(res.Videos))
	for i, video := range res.Videos {
		videos[i] = SmallVideo2Response(video)
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取点赞列表成功", refreshedToken, gin.H{
		"videos":   videos,
		"nextTime": res.NextTime,
	})
}

func GetCollectList(c *gin.Context) {
	latestTime := util.String2Int64(c.Query("latestTime"))
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	uid := util.String2Int64(c.Query("uid"))
	currentUserId := c.GetInt("userId")
	if uid == 0 {
		uid = int64(currentUserId)
	}
	res, err := global.InteractionServiceClient.GetCollectList(c, &interactionProto.GetCollectListRequest{
		UserId:     uid,
		LatestTime: latestTime,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	videos := make([]gin.H, len(res.Videos))
	for i, video := range res.Videos {
		videos[i] = SmallVideo2Response(video)
	}
	var refreshedToken interface{}
	if currentUserId != 0 {
		token := c.GetString("token")
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取收藏列表成功", refreshedToken, gin.H{
		"videos":   videos,
		"nextTime": res.NextTime,
	})
}

func Comment(c *gin.Context) {
	commentForm := form.CommentForm{}
	if err := c.ShouldBind(&commentForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	uid := c.GetInt("userId")
	_, err := global.InteractionServiceClient.Comment(c, &interactionProto.CommentRequest{
		UserId:   int64(uid),
		VideoId:  util.String2Int64(commentForm.VideoId),
		Content:  commentForm.Content,
		ParentId: int64(commentForm.ParentId),
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "评论成功", refreshedToken, nil)
}

func DeleteComment(c *gin.Context) {
	cid := util.String2Int64(c.Query("cid"))
	if cid == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "cid不能为空", nil, nil)
		return
	}
	uid := c.GetInt("userId")
	_, err := global.InteractionServiceClient.DeleteComment(c, &interactionProto.DeleteCommentRequest{
		UserId:    int64(uid),
		CommentId: cid,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "删除评论成功", refreshedToken, nil)
}

func GetCommentList(c *gin.Context) {
	vid := util.String2Int64(c.Query("vid"))
	if vid == 0 {
		HandleHttpResponse(c, http.StatusBadRequest, "vid不能为空", nil, nil)
		return
	}
	res, err := global.InteractionServiceClient.GetCommentList(c, &interactionProto.GetCommentListRequest{
		VideoId: vid,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	var refreshedToken interface{}
	token := c.GetString("token")
	if token != "" {
		j := middleware.NewJWT()
		refreshedToken, _ = j.RefreshToken(token)
	}
	HandleHttpResponse(c, http.StatusOK, "获取评论列表成功", refreshedToken, res.Comments)
}

func SmallVideo2Response(video *interactionProto.SmallVideo) gin.H {
	return gin.H{
		"vid":         strconv.FormatInt(video.Id, 10),
		"coverUrl":    video.CoverUrl,
		"playUrl":     video.PlayUrl,
		"likeCount":   video.LikeCount,
		"description": video.Description,
		"topics":      video.Topics,
	}
}
