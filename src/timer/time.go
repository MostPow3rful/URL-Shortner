package timer

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func SortDate(_date string) (year, month, day, hour, minute, second int) {
	var (
		expireDate [3]int   = [3]int{}
		expireTime [3]int   = [3]int{}
		nowTime    []string = []string{}
	)
	nowTime = strings.Split(_date, " ")[0:2]

	for key, value := range strings.Split(nowTime[0], "-") {
		expireDate[key], _ = strconv.Atoi(value)
	}

	for key, value := range strings.Split(nowTime[1], ":") {
		expireTime[key], _ = strconv.Atoi(strings.Split(value, ".")[0])
	}

	return expireDate[0], expireDate[1], expireDate[2], expireTime[0], expireTime[1], expireTime[2]
}

func SetExpire(_hours string) string {
	if _hours == "0" {
		year, month, day, hour, minute, second := SortDate(time.Now().Add(time.Hour * 999999).String())
		return fmt.Sprintf("%d-%d-%d %d.%d.%d", year, month, day, hour, minute, second)
	}
	hours, _ := strconv.Atoi(_hours)
	year, month, day, hour, minute, second := SortDate(time.Now().Add(time.Hour * time.Duration(hours)).String())
	return fmt.Sprintf("%d-%d-%d %d.%d.%d", year, month, day, hour, minute, second)
}

func CheckExpire(_expire string) bool {
	nowYear, nowMonth, nowDay, nowHour, nowMinute, nowSecond := SortDate(time.Now().String())
	year, month, day, hour, minute, second := SortDate(_expire)

	if (nowYear > year) || (nowMonth > month) || (nowDay > day) {
		return false
	}

	if (nowYear == year) && (nowMonth == month) && (nowDay == day) {
		if nowHour > hour {
			return false
		} else if (nowHour == hour) && (nowMinute > minute) {
			return false
		} else if (nowHour == hour) && (nowMinute == minute) && (nowSecond == second) {
			return false
		}
	}

	return true
}
