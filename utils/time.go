package utils

import (
	"strings"

	"github.com/nosixtools/solarlunar"
)

func GetChineseDate(solarDate string) string {
	chineseDate := solarlunar.SolarToSimpleLuanr(solarDate)
	chineseDate = strings.ReplaceAll(chineseDate, "年", "-")
	chineseDate = strings.ReplaceAll(chineseDate, "月", "-")
	chineseDate = strings.ReplaceAll(chineseDate, "日", "")
	return chineseDate
}
