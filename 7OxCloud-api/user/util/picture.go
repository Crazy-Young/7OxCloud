package util

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/global"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func Upload(file multipart.File, header *multipart.FileHeader) (picUrl string, err error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}
	key := fmt.Sprintf("%s%s", GenerateUUID(), filepath.Ext(header.Filename))
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", global.ServerConfig.QiNiuYun.Bucket, key),
	}
	mac := qbox.NewMac(global.ServerConfig.QiNiuYun.AccessKey, global.ServerConfig.QiNiuYun.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err = uploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), header.Size, &putExtra)
	if err != nil {
		return
	}
	picUrl = fmt.Sprintf("%s/%s", global.ServerConfig.QiNiuYun.Domain, ret.Key)
	return
}

func Delete(picUrl string) (err error) {
	mac := qbox.NewMac(global.ServerConfig.QiNiuYun.AccessKey, global.ServerConfig.QiNiuYun.SecretKey)
	bucketManager := storage.NewBucketManager(mac, nil)
	key := filepath.Base(picUrl)
	err = bucketManager.Delete(global.ServerConfig.QiNiuYun.Bucket, key)
	return
}
