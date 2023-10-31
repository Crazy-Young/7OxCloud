package api

import (
	"errors"
	"net/http"
	"strings"

	videoProto "github.com/palp1tate/7OxCloud/proto/video"

	userProto "github.com/palp1tate/7OxCloud/proto/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/global"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

func HandleGrpcErrorToHttp(c *gin.Context, err error) {
	if e, ok := status.FromError(err); ok {
		HandleHttpResponse(c, http.StatusBadRequest, e.Message(), nil, nil)
		return
	}
}

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(c *gin.Context, err error) {
	var errs validator.ValidationErrors
	if ok := errors.As(err, &errs); !ok {
		zap.S().Errorw("验证器参数校验失败", "err", err.Error())
		HandleHttpResponse(c, http.StatusInternalServerError, err.Error(), nil, nil)
	} else {
		zap.S().Info("参数格式错误")
		HandleHttpResponse(c, http.StatusBadRequest, removeTopStruct(errs.Translate(global.Translator)), nil, nil)
	}
	return
}

func HandleHttpResponse(c *gin.Context, code int, msg interface{}, token interface{}, data interface{}) {
	h := gin.H{
		"code": code,
		"msg":  msg,
	}
	if token != nil {
		h["token"] = token
	}
	if data != nil {
		h["data"] = data
	}
	c.JSON(code, h)
	return
}

func UserToResponse(user *userProto.User) gin.H {
	return gin.H{
		"uid":          user.Id,
		"username":     user.Username,
		"location":     user.Location,
		"age":          user.Age,
		"avatar":       user.Avatar,
		"gender":       user.Gender,
		"signature":    user.Signature,
		"isFollow":     user.IsFollow,
		"followCount":  user.FollowCount,
		"fanCount":     user.FanCount,
		"likeCount":    user.LikeCount,
		"likedCount":   user.LikedCount,
		"workCount":    user.WorkCount,
		"collectCount": user.CollectCount,
	}
}

func MiniUserToResponse(user *videoProto.User) gin.H {
	return gin.H{
		"uid":      user.Id,
		"username": user.Username,
		"avatar":   user.Avatar,
		"location": user.Location,
		"isFollow": user.IsFollow,
	}
}

func VideoToResponse(video *videoProto.Video) gin.H {
	return gin.H{
		"vid":          video.Id,
		"description":  video.Description,
		"playUrl":      video.PlayUrl,
		"coverUrl":     video.CoverUrl,
		"likeCount":    video.LikeCount,
		"commentCount": video.CommentCount,
		"collectCount": video.CollectCount,
		"isLike":       video.IsLike,
		"author":       MiniUserToResponse(video.Author),
		"topics":       video.Topics,
		"publishTime":  video.PublishTime,
	}
}

func MiniVideoToResponse(video *videoProto.MiniVideo) gin.H {
	return gin.H{
		"vid":       video.Id,
		"coverUrl":  video.CoverUrl,
		"likeCount": video.LikeCount,
	}
}
