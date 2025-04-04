# 阿里云短信发送/查询（单条）

## 功能
官方的SDK看起来很复杂，功能很多；我就把最常用的单条短信收发给封装了一下。
更简单易懂，调用、传参、返回结果！

## 使用说明

一、安装

`go get "github.com/luji3120346/aliSms"`

二、引用aliSMsApi包

`import "github.com/luji3120346/aliSms/aliSmsApi"`

三、创建类型实例 AliSms

`client := aliSmsApi.AliSms{}`

四、实例方法

>`CreatClient(accessKeyId,accessKeySecret)  //创建客户端`

>`SendMsm(aliSmsApi.SendSmsConfig)(aliSmsApi.SendSmsReceipt,error)  //发送短信`

>`QuerySms(aliSmsApi.SendSmsReceipt)(string,error) //查询短信`

### 使用示例：

```go
func main() {

    //新建客户端实例
    client := aliSmsApi.AliSms{}
    err := client.CreatClient("这里填阿里云的accessKeyId", "这里填阿里云的accessKeySecret")
    if err != nil {
        log.Println(err)
    }

    //发送短信
    receipt, err := client.SendSms(aliSmsApi.SendSmsConfig{
        PhoneNumbers:  "13111112222", //手机号
        SignName:      "USTOOLS网站", //签名名称
        TemplateCode:  "SMS_300560153", //短信模板代码
        TemplateParam: "{'code':'9941'}", //JSON字符串，模板中的变量设置
    })
    if err != nil {
        log.Println(err)
    }
    log.Println(receipt)

    //查询短信
    content,err :=client.QuerySms(aliSmsApi.SendSmsReceipt{
        PhoneNumbers:  "13111112222", //手机号
        BizId:      "12312419023812093^1", //发送短信时，服务器返回的BizId
        Date:  "20250404", //发送短信的日期
    })
    if err != nil {
        log.Println(err)
    }
    log.Println(content)
}
```

## 与我联系

有啥问题请留言联系
>邮箱 luji3120346@126.com