package service

import (
	"birth/consts"
	"birth/model"
	"birth/utils"
	"io/ioutil"
	"strings"
	"time"
)

var BirthDayServiceIstance *BirthDayService = new(BirthDayService)

type BirthDayService struct {
	birthDays []*model.BirthDay
}

func (b *BirthDayService) Init(configPath string) {
	b.loadConfig(configPath)
	// res, _ := json.Marshal(b.birthDays)
	// fmt.Print(string(res))
}
func (b *BirthDayService) loadConfig(configPath string) {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	birthdayItems := strings.Split(string(content), "\n")
	var birthDays []*model.BirthDay
	for _, item := range birthdayItems {
		birthDay := model.NewBirthDay(item)
		if len(birthDay.ChineseDate)+len(birthDay.SolarDate) != 5 {
			panic("格式错误")
		}
		birthDays = append(birthDays, birthDay)
	}

	b.birthDays = birthDays
}

func (b *BirthDayService) GetTodayBirthDays() []*model.BirthDay {
	soloarNow := time.Now().In(consts.CSTZone).Format("2006-01-02")
	chineseNow := utils.GetChineseDate(soloarNow)
	soloarNowWithoutYear := soloarNow[5:]
	chineseNowWithoutYear := chineseNow[5:]
	todayBirthDay := make([]*model.BirthDay, 0)
	for _, birthDay := range b.birthDays {
		if birthDay.IsSolarDate && birthDay.SolarDate == soloarNowWithoutYear {
			todayBirthDay = append(todayBirthDay, birthDay)
		} else if birthDay.ChineseDate == chineseNowWithoutYear {
			todayBirthDay = append(todayBirthDay, birthDay)
		}
	}
	return todayBirthDay
}
