package sms_util

import (
	"CapPrice/logging"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/spf13/viper"
)

func SendHandler(cellphone, code string) error {
	client, err := dysmsapi.NewClientWithAccessKey(
		viper.GetString("aliyun_sms.region"),
		viper.GetString("aliyun_sms.ak"),
		viper.GetString("aliyun_sms.aks"),
	)

	if err != nil {
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = cellphone
	request.SignName = viper.GetString("aliyun_sms.sign_name")
	request.TemplateCode = viper.GetString("aliyun_sms.template_code")
	request.TemplateParam = fmt.Sprintf("{\"code\": \"%s\"}", code)

	if _, err = client.SendSms(request); err != nil {
		fmt.Print(err.Error())
		logging.STDError("发送验证码失败: %v", err)
		return err
	}
	return nil
}
