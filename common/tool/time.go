package tool

import (
	"strings"
	"time"
)

var FixedZone = time.FixedZone("CST", 8*3600)

const (
	DATE_LAYOUT         = "2006-01-02"
	DATETIME_LAYOUT     = "2006-01-02 15:04:05"
	UTC_DATETIME_LAYOUT = "2006-01-02T15:04:05Z0700"
)

func TimeConvertToDate(time *time.Time) string {
	date := time.In(FixedZone).Format("2006-01-02")
	if date == "0001-01-01" {
		return ""
	}
	return date
}
func TimeConvertToDatetime(time *time.Time) string {
	dateTime := time.In(FixedZone).Format("2006-01-02 15:04:05")
	if dateTime == "0001-01-01 00:00:00" {
		return ""
	}
	return dateTime
}

func DateConvertToTime(date string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", date, FixedZone)
}
func DatetimeConvertToTime(datetime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", datetime, FixedZone)
}
func UTCDatetimeConvertToTime(datetime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02T15:04:05Z0700", datetime, FixedZone)
}
func GetDayNum(startDate string, endDate string) int {
	startTime, _ := DateConvertToTime(startDate)
	endTime, _ := DateConvertToTime(endDate)
	return GetDayCount(startTime, endTime)
}
func GetDayCount(startDate time.Time, endDate time.Time) int {
	return int(endDate.Sub(startDate).Hours()/24) + 1
}

/**
将格式为yyyy-MM-dd hh-MM-ss转化为yyyy-MM-dd
 */
func DateTimeConvertDate(dateTime string) string {
	s := strings.Split(dateTime, " ")
	return s[0]
}

/**
获取特定x天数前的日期（不包含非工作日）
 */
func GetLastDayWithNonWorkDay(nonWorkDay string) time.Time {
	date := time.Now()
	flag := 5
	var num int
	for flag != 0 {
		date = date.AddDate(0,0,-1)
		if strings.Contains(nonWorkDay, TimeConvertToDate(&date)) {
			num++
		}else {
			num++
			flag--
		}
	}
	curDate := time.Now()
	date = curDate.AddDate(0, 0, -num)
	return date
}

/**
获取当天上一周周一的日期
 */
func GetLastWeekDate() time.Time {
	date := time.Now()
	week := date.Weekday().String()
	if week == "Monday" {
		date = date.AddDate(0, 0, -7)
	}else if week == "Tuesday" {
		date = date.AddDate(0, 0, -8)
	}else if week == "Wednesday" {
		date = date.AddDate(0, 0, -9)
	}else if week == "Thursday" {
		date = date.AddDate(0, 0, -10)
	}else if week == "Friday" {
		date = date.AddDate(0, 0, -11)
	}else if week == "Saturday" {
		date = date.AddDate(0, 0, -12)
	}else {
		date = date.AddDate(0, 0, -13)
	}
	return date
}