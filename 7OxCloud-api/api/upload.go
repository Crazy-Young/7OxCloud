package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/form"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/middleware"
	"github.com/palp1tate/7OxCloud/util"
	"go.uber.org/zap"
)

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		zap.S().Info("文件上传失败:", err)
		HandleHttpResponse(c, http.StatusBadRequest, "文件上传失败", nil, nil)
		return
	}
	url, err := util.Upload(file, header)
	if err != nil {
		zap.S().Info("文件上传失败:", err.Error())
		HandleHttpResponse(c, http.StatusBadRequest, "文件上传失败", nil, nil)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "文件上传成功", refreshedToken, url)
	return
}

func DeleteFile(c *gin.Context) {
	urlForm := form.UrlForm{}
	if err := c.ShouldBind(&urlForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	url := urlForm.Url
	err := util.Delete(url)
	if err != nil {
		zap.S().Info("文件删除失败:", err.Error())
		HandleHttpResponse(c, http.StatusBadRequest, "文件删除失败", nil, nil)
		return
	}
	token := c.GetString("token")
	j := middleware.NewJWT()
	refreshedToken, _ := j.RefreshToken(token)
	HandleHttpResponse(c, http.StatusOK, "文件删除成功", refreshedToken, nil)
	return
}
