package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/consts"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/form"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/middleware"
	"github.com/palp1tate/7OxCloud/proto/user"
	"github.com/palp1tate/7OxCloud/util"
	"go.uber.org/zap"
)

func Register(c *gin.Context) {
	registerForm := form.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	value, err := global.RedisClient.Get(context.Background(), registerForm.Mobile).Result()
	if err == redis.Nil {
		HandleHttpResponse(c, http.StatusBadRequest, "验证码已过期", nil, nil)
		return
	} else {
		if value != registerForm.Captcha {
			HandleHttpResponse(c, http.StatusBadRequest, "验证码错误", nil, nil)
			return
		}
		global.RedisClient.Del(context.Background(), registerForm.Mobile)
	}
	res, err := global.UserServiceClient.Register(context.Background(), &userProto.RegisterRequest{
		Mobile:   registerForm.Mobile,
		Password: registerForm.Password,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	claims := middleware.CustomClaims{
		ID: int(res.Id),
	}
	j := middleware.NewJWT()
	token, err := j.CreateToken(claims)
	if err != nil {
		HandleHttpResponse(c, http.StatusInternalServerError, "生成token失败", nil, nil)
		return
	}
	HandleHttpResponse(c, http.StatusOK, "注册成功", token, nil)
	return
}

func Login(c *gin.Context) {
	loginForm := form.LoginForm{}
	if err := c.ShouldBind(&loginForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	var token string
	switch loginForm.Type {
	case consts.LoginByPassword:
		if ok := util.Store.Verify(loginForm.CaptchaId, loginForm.Captcha, true); !ok {
			HandleHttpResponse(c, http.StatusBadRequest, "验证码错误", nil, nil)
			return
		}
		res, err := global.UserServiceClient.LoginByPassword(context.Background(), &userProto.LoginByPasswordRequest{
			Mobile:   loginForm.Mobile,
			Password: loginForm.Password,
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		claims := middleware.CustomClaims{
			ID: int(res.Id),
		}
		j := middleware.NewJWT()
		token, err = j.CreateToken(claims)
		if err != nil {
			HandleHttpResponse(c, http.StatusInternalServerError, "生成token失败", nil, nil)
			return
		}
	case consts.LoginBySMS:
		value, err := global.RedisClient.Get(context.Background(), loginForm.Mobile).Result()
		if err == redis.Nil {
			HandleHttpResponse(c, http.StatusBadRequest, "验证码已过期", nil, nil)
			return
		} else {
			if value != loginForm.Captcha {
				HandleHttpResponse(c, http.StatusBadRequest, "验证码错误", nil, nil)
				return
			}
			global.RedisClient.Del(context.Background(), loginForm.Mobile)
		}
		res, err := global.UserServiceClient.LoginBySMS(context.Background(), &userProto.LoginBySMSRequest{
			Mobile: loginForm.Mobile,
		})
		if err != nil {
			HandleGrpcErrorToHttp(c, err)
			return
		}
		claims := middleware.CustomClaims{
			ID: int(res.Id),
		}
		j := middleware.NewJWT()
		token, err = j.CreateToken(claims)
		if err != nil {
			HandleHttpResponse(c, http.StatusInternalServerError, "生成token失败", nil, nil)
			return
		}
	}
	HandleHttpResponse(c, http.StatusOK, "登录成功", token, nil)
	return
}

func GetUser(c *gin.Context) {
	currentUserId := int64(c.GetInt("userId"))
	uid := util.String2Int64(c.Query("uid"))
	if uid == 0 {
		uid = currentUserId
	}
	res, err := global.UserServiceClient.GetUser(context.Background(), &userProto.GetUserRequest{
		Id:            uid,
		CurrentUserId: currentUserId,
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
	HandleHttpResponse(c, http.StatusOK, "获取用户信息成功", refreshedToken, UserToResponse(res.User))
	return
}

func UpdateUser(c *gin.Context) {
	updateUserForm := form.UpdateUserForm{}
	if err := c.ShouldBind(&updateUserForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	uid := int64(c.GetInt("userId"))
	_, err := global.UserServiceClient.UpdateUser(context.Background(), &userProto.UpdateUserRequest{
		Id:        uid,
		Username:  updateUserForm.Username,
		Location:  updateUserForm.Location,
		Age:       int64(updateUserForm.Age),
		Avatar:    updateUserForm.Avatar,
		Gender:    util.String2Int64(updateUserForm.Gender),
		Signature: updateUserForm.Signature,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "更新用户信息成功", refreshedToken, nil)
	return
}

func ResetPassword(c *gin.Context) {
	resetPasswordForm := form.ResetPasswordForm{}
	if err := c.ShouldBind(&resetPasswordForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	value, err := global.RedisClient.Get(context.Background(), resetPasswordForm.Mobile).Result()
	if err == redis.Nil {
		zap.S().Info("验证码已过期")
		HandleHttpResponse(c, http.StatusBadRequest, "验证码已过期", nil, nil)
		return
	} else {
		if value != resetPasswordForm.Captcha {
			zap.S().Info("验证码错误")
			HandleHttpResponse(c, http.StatusBadRequest, "验证码错误", nil, nil)
			return
		}
		global.RedisClient.Del(context.Background(), resetPasswordForm.Mobile)
	}
	_, err = global.UserServiceClient.ResetPassword(context.Background(), &userProto.ResetPasswordRequest{
		Mobile:   resetPasswordForm.Mobile,
		Password: resetPasswordForm.Password,
	})
	if err != nil {
		HandleGrpcErrorToHttp(c, err)
		return
	}
	HandleHttpResponse(c, http.StatusOK, "修改密码成功", nil, nil)
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
