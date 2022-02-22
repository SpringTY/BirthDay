package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"gopkg.in/gomail.v2"
)

var SMTPMailServiceIstance *SMTPMailService = new(SMTPMailService)

type SMTPMailService struct {
	smtpServer    string
	smtpPort      int
	smtpUserName  string
	smtpAuthToken string
}

func (s *SMTPMailService) Init(configPath string) {
	configMap := make(map[string]string)
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &configMap)
	if err != nil {
		panic(err)
	}
	stmpPort, _ := strconv.Atoi(configMap["stmp_port"])
	s.smtpServer = configMap["stmp_server"]
	s.smtpPort = stmpPort
	s.smtpUserName = configMap["stmp_user_name"]
	s.smtpAuthToken = configMap["stmp_auth_token"]
}
func (s *SMTPMailService) SendHtmlMail(receiverMail, subject, body string) error {
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", s.smtpUserName)
	//接收人
	m.SetHeader("To", receiverMail)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", subject)
	//内容
	m.SetBody("text/html", body)

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer(s.smtpServer, s.smtpPort, s.smtpUserName, s.smtpAuthToken)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		return err
	}
	fmt.Printf("send mail success\n")
	return nil
}
