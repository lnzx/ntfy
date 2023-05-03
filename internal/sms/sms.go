package sms

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"log"
	"os"
)

var appId string
var signName string
var templateId string

var client *sms.Client

func init() {
	appId = os.Getenv("APP_ID")
	signName = os.Getenv("SIGN_NAME")
	templateId = os.Getenv("TEMPLATE_ID")

	credential := common.NewCredential(os.Getenv("SECRET_ID"), os.Getenv("SECRET_KEY"))

	cpf := profile.NewClientProfile()

	var err error
	client, err = sms.NewClient(credential, regions.Guangzhou, cpf)
	if err != nil {
		log.Fatalln("init sms:", err)
	}
}

func Send(mobile string, event string) error {
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = &appId
	request.SignName = &signName
	request.TemplateId = &templateId

	request.PhoneNumberSet = []*string{&mobile}
	request.TemplateParamSet = []*string{&event}

	rsp, err := client.SendSms(request)
	if err != nil {
		return err
	}
	log.Println(rsp.ToJsonString())
	return nil
}
