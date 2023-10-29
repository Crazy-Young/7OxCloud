package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/form"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/middleware"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/util"
	"go.uber.org/zap"
)

func UploadPicture(c *gin.Context) {
	file, header, err := c.Request.FormFile("picture")
	if err != nil {
		zap.S().Info("图片上传失败:", err)
		HandleHttpResponse(c, http.StatusBadRequest, "图片上传失败", nil, nil)
		return
	}
	picUrl, err := util.Upload(file, header)
	if err != nil {
		zap.S().Info("图片上传失败:", err.Error())
		HandleHttpResponse(c, http.StatusBadRequest, "图片上传失败", nil, nil)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "图片上传成功", refreshedToken, gin.H{
		"picUrl": picUrl,
	})
	return
}

func DeletePicture(c *gin.Context) {
	picUrlForm := form.PicUrlForm{}
	if err := c.ShouldBind(&picUrlForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	picUrl := picUrlForm.PicUrl
	err := util.Delete(picUrl)
	if err != nil {
		zap.S().Info("图片删除失败:", err.Error())
		HandleHttpResponse(c, http.StatusBadRequest, "图片删除失败", nil, nil)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "图片删除成功", refreshedToken, nil)
	return
}
