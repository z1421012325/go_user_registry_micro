package utils

import (
	"github.com/go-gomail/gomail"
	"fmt"
	"log"
)

var (
	user = `xxxxxx@qq.com` 	   // 发送邮箱：账号
	pass = `xxxxxx`            // 发送邮箱：密码（qq邮箱：密码填授权码）
	host = `smtp.qq.com`       // 发送邮箱：服务器地址
	port = 25                  // 发送邮箱：端口（默认端口：465，QQ邮箱端口：25）
)


func SendEmail(mailTo,title,body string){
	m := gomail.NewMessage()
	m.SetHeader(`From`, user)
	m.SetHeader(`To`, mailTo)
	m.SetHeader(`Subject`, title)
	m.SetBody(`text/html`, body)

	err := gomail.NewDialer(host, port, user, pass).DialAndSend(m)
	if err != nil {
		log.Println("send email is err : ",err)
		// 异常处理 数据库保存 发送邮件or短信报错
	}
}

// todo 测试能不能发送
func SendRegistryEmail(mailTo,Strcode string){
	body := fmt.Sprintf("xxx 注册账号的验证码为 %s ,有效时间 6 分钟",Strcode)
	SendEmail(mailTo,"xxx 注册账号",body)
}

func SendResetPassWordEmail(mailTo,Strcode string){
	body := fmt.Sprintf("xxx 重置账号密码的验证码为 %s ,有效时间 6 分钟",Strcode)
	SendEmail(mailTo,"xxx 密码重置",body)
}

// .... sendxxxxEmail ....