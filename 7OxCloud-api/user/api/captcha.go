package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/util"
	"go.uber.org/zap"
)

func GetCaptcha(ctx *gin.Context) {
	id, b64s, err := util.GeneratePicCaptcha()
	if err != nil {
		zap.S().Info("生成验证码错误: ", err.Error())
		HandleHttpResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
	}
	HandleHttpResponse(ctx, http.StatusOK, "获取图片验证码成功", nil, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}
