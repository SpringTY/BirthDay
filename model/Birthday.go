package model

import (
	"strings"
)

type BirthDay struct {
	Name        string `json:"name"`
	SolarDate   string `json:"solar_date"`
	ChineseDate string `json:"chinese_date"`
	IsSolarDate bool   `json:"is_solar_date"`
}

func NewBirthDay(configItem string) *BirthDay {
	// fmt.Println(configItem)
	birthdayInfo := strings.Fields(configItem)
	// fmt.Println(len(birthdayInfo))
	birthDay := new(BirthDay)
	birthDay.Name = birthdayInfo[0]
	if len(birthdayInfo) == 2 {
		birthDay.IsSolarDate = true
		birthDay.SolarDate = birthdayInfo[1]
	} else {
		birthDay.IsSolarDate = false
		birthDay.ChineseDate = birthdayInfo[2]
	}
	return birthDay
}
