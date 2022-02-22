package main

import (
	"birth/handler"
	"birth/service"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	birthDayConfig := "./conf/birth.txt"
	mailConfig := "./conf/mail.json"
	service.BirthDayServiceIstance.Init(birthDayConfig)
	service.SMTPMailServiceIstance.Init(mailConfig)
	handler.SubjectBirthDay("xxxx@qq.com")

	c := cron.New()

	c.AddFunc("0 0 12 * * ?", handler.WrapSubjectBirthDay)

	c.Start()
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println("stil on")
	}
}
