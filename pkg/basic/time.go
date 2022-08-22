package basic

import "time"

// 获取当天时间戳
func GetDateUnix() int64 {
	currentTime := time.Now()
	currentDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), Zero, Zero, Zero, Zero, currentTime.Location())
	return currentDay.Unix()
}
