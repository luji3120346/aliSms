package aliSmsApi

import (
	"fmt"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// 这是测试二班
type AliSms struct {
	Client *dysmsapi20170525.Client
}

type SendSmsConfig struct {
	PhoneNumbers  string `json:"phone_numbers"`  // 手机号码
	SignName      string `json:"sign_name"`      // 短信签名
	TemplateCode  string `json:"template_code"`  // 短信模板ID
	TemplateParam string `json:"template_param"` // 短信模板变量替换JSON串
}

type SendSmsReceipt struct {
	PhoneNumbers string `json:"phone_numbers"` // 手机号码
	BizId        string `json:"biz_id"`        // 短信发送流水号
	Date         string `json:"date"`          // 短信发送时间YYMMDD
}

// CreatClient 创建阿里云短信客户端
// keyId 和 keySecret 可以在阿里云控制台获取
func (c *AliSms) CreatClient(keyId string, keySecret string) error {
	config := &openapi.Config{
		AccessKeyId:     tea.String(keyId),
		AccessKeySecret: tea.String(keySecret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}

	var err error
	c.Client, err = dysmsapi20170525.NewClient(config)
	if err != nil {
		return fmt.Errorf("创建阿里云短信客户端失败：%v", err)
	}
	return nil
}

// SendMsm 发送短信
func (c *AliSms) SendMsm(info SendSmsConfig) (SendSmsReceipt, error) {
	sendSmsReceipt := SendSmsReceipt{}
	if c.Client == nil {
		return sendSmsReceipt, fmt.Errorf("短信客户端未初始化，请先调用 CreatClient 方法")
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(info.PhoneNumbers),
		SignName:      tea.String(info.SignName),
		TemplateCode:  tea.String(info.TemplateCode),
		TemplateParam: tea.String(info.TemplateParam),
	}

	result, err := c.Client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
	if err != nil {
		return sendSmsReceipt, fmt.Errorf("发送短信失败：%v;请求参数为：%v", err, info)
	}
	if *result.Body.Code != "OK" {
		return sendSmsReceipt, fmt.Errorf("发送短信失败：%v;请求参数为：%v", *result.Body.Message, info)
	}
	sendSmsReceipt.PhoneNumbers = info.PhoneNumbers
	sendSmsReceipt.BizId = *result.Body.BizId
	sendSmsReceipt.Date = time.Now().Format("20060102")

	return sendSmsReceipt, nil
}

// QuerySms 查询短信发送详情
func (c *AliSms) QuerySms(info SendSmsReceipt) (string, error) {
	var content string
	if c.Client == nil {
		return content, fmt.Errorf("短信客户端未初始化，请先调用 CreatClient 方法")
	}

	querySendDetailsRequest := &dysmsapi20170525.QuerySendDetailsRequest{
		PhoneNumber: tea.String(info.PhoneNumbers),
		SendDate:    tea.String(info.Date),
		PageSize:    tea.Int64(1),
		CurrentPage: tea.Int64(1),
		BizId:       tea.String(info.BizId),
	}

	result, err := c.Client.QuerySendDetailsWithOptions(querySendDetailsRequest, &util.RuntimeOptions{})
	if err != nil {
		return content, fmt.Errorf("查询短信发送详情失败：%v;请求参数为：%v", err, info)
	}
	if *result.Body.Code != "OK" {
		return content, fmt.Errorf("查询短信发送详情失败：%v;请求参数为：%v", *result.Body.Message, info)
	}

	contentList := result.Body.SmsSendDetailDTOs.SmsSendDetailDTO
	if len(contentList) == 0 {
		return content, fmt.Errorf("没有找到相关短信记录，查询参数为：%v", info)
	}

	content = *contentList[0].Content
	return content, nil
}
