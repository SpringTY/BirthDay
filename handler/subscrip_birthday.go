package handler

import (
	"birth/consts"
	"birth/model"
	"birth/service"
	"fmt"
	"time"
)

func WrapSubjectBirthDay() {
	SubjectBirthDay("xxxx@qq.com")
	//SubjectBirthDay("370286558@qq.com")

}

func SubjectBirthDay(email string) {
	todayBirthdays := service.BirthDayServiceIstance.GetTodayBirthDays()
	if len(todayBirthdays) == 0 {
		return
	}
	body := buildSubjectBirthDayEmailBody(todayBirthdays)
	subject := "insis生日提醒"
	service.SMTPMailServiceIstance.SendHtmlMail(email, subject, body)
	//make body
	// Insis生日订阅者：\n\t您好！\n\t今天是: %s.共有以下几位同学今天生日:%s \n\s
}
func buildSubjectBirthDayEmailBody(birthdays []*model.BirthDay) string {
	rawFormatter := " Insis生日订阅者：<div/>您好！<div/>今天是: %s.共有以下几位同学今天生日:%s <div/> "
	names := ""
	for _, birthday := range birthdays {
		names += birthday.Name + ", "
	}
	names = names[:len(names)-2]
	soloarNow := time.Now().In(consts.CSTZone).Format("2006-01-02")
	return fmt.Sprintf(rawFormatter, soloarNow, names)
}
