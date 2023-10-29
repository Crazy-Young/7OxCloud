package util

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	aliutil "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/global"
)

func GenerateSmsCaptcha(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789"
	var result strings.Builder
	for i := 0; i < length; i++ {
		result.WriteByte(charset[r.Intn(len(charset))])
	}
	return result.String()
}

func CreateClient(accessKeyId *string, accessKeySecret *string) (client *dysmsapi20170525.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	client = &dysmsapi20170525.Client{}
	client, err = dysmsapi20170525.NewClient(config)
	return client, err
}

func SendSms(mobile string) (captcha string, err error) {
	accessKeyId := global.ServerConfig.AliSms.AccessKeyId
	accessKeySecret := global.ServerConfig.AliSms.AccessKeySecret
	signName := global.ServerConfig.AliSms.SignName
	templateCode := global.ServerConfig.AliSms.TemplateCode
	client, err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		return "", err
	}
	captcha = GenerateSmsCaptcha(6)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(mobile),
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(templateCode),
		TemplateParam: tea.String(fmt.Sprintf(`{"code":"`+"%s"+`"}`, captcha)),
	}
	defer func() {
		if r := tea.Recover(recover()); r != nil {
			err = r
		}
	}()
	runtime := &aliutil.RuntimeOptions{}
	result, err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil {
		return
	}
	if *result.Body.Code != "OK" {
		err = errors.New(result.String())
		return
	}
	return
}
