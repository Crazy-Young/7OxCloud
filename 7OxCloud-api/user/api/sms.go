package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/consts"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/form"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/proto"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/util"
	"go.uber.org/zap"
)

func SendSms(c *gin.Context) {
	sendSmsForm := form.SendSmsForm{}
	if err := c.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	switch sendSmsForm.Type {
	case consts.Register:
		if res, err := global.UserServiceClient.CheckMobile(context.Background(), &proto.CheckMobileRequest{
			Mobile: sendSmsForm.Mobile,
		}); err != nil {
			zap.S().Info("服务器内部错误")
			HandleGrpcErrorToHttp(c, err)
			return
		} else if res.Exist {
			zap.S().Info("该手机号已经被注册")
			HandleHttpResponse(c, http.StatusBadRequest, "该手机号已经被注册", nil, nil)
			return
		}
	default:
		if res, err := global.UserServiceClient.CheckMobile(context.Background(), &proto.CheckMobileRequest{
			Mobile: sendSmsForm.Mobile,
		}); err != nil {
			zap.S().Info("服务器内部错误")
			HandleGrpcErrorToHttp(c, err)
			return
		} else if !res.Exist {
			zap.S().Info("该手机号未注册")
			HandleHttpResponse(c, http.StatusBadRequest, "该手机号未注册", nil, nil)
			return
		}
	}
	smsCode, err := util.SendSms(sendSmsForm.Mobile)
	if err != nil {
		zap.S().Info("发送短信失败")
		HandleHttpResponse(c, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	global.RedisClient.Set(context.Background(), sendSmsForm.Mobile, smsCode,
		time.Duration(global.ServerConfig.Redis.Expiration)*time.Minute)
	HandleHttpResponse(c, http.StatusOK, "短信验证码发送成功", nil, nil)
	return
}
