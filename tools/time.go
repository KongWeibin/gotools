package tools

import (
	"fmt"
	"strconv"
	"time"
)

//时间相关的一些常值
const (
	MonthsPerYear            = 12
	SecondsPerHour           = 3600
	NanoSecondPerMillisecond = 1000000
	FormatLayout             = "2006-01-02 15:04:05" //time 包里固定格式化字符串之一
)

//GetNowUnix 获取当前时间时间戳-秒
func GetNowUnix() int64 {
	return time.Now().Unix()
}

//GetNowUnixMilli 获取当前时间时间戳-毫秒
func GetNowUnixMilli() int64 {
	return time.Now().UnixNano() / NanoSecondPerMillisecond
}

//GetNowUnixNano 获取当前时间时间戳-纳秒
func GetNowUnixNano() int64 {
	return time.Now().UnixNano()
}

//GetNowFarmat 获取当前时间，格式化"2006-01-02 15:04:05"输出
func GetNowFarmat() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

// GetDateAfterNDays 获取n天后日期字符串，格式:20181212,ndays 可正（未来）可负（过去）
func GetDateAfterNDays(ndays int) string {
	destDay := time.Now().Add(time.Duration(ndays*86400) * time.Second)
	y, m, d := destDay.Date()
	return fmt.Sprintf("%d%02d%02d", y, m, d)
}

// GetTodayStartTimeStamp 获取当天零点时间戳
func GetTodayStartTimeStamp() int64 {
	tStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", tStr)
	ti := t.Unix()
	//格式化后是上午08：00：00
	return ti - 8*SecondsPerHour
}

//GetMonthAfterNMonths 获取N月后的月份日期 格式:201812,nmonths 可正（未来）可负（过去）
func GetMonthAfterNMonths(nmonths int) string {
	yi := time.Now().Year()
	mi := int(time.Now().Month())

	if Abs(int64(mi+nmonths)) > MonthsPerYear && (mi+nmonths)%MonthsPerYear == 0 {
		yi = yi + (mi+nmonths)/MonthsPerYear - 1
		mi = 12 //月份都是12月
	} else if Abs(int64(mi+nmonths)) > MonthsPerYear && (mi+nmonths)%MonthsPerYear != 0 {
		yi = yi + (mi+nmonths)/MonthsPerYear
		mi = (mi + nmonths) % MonthsPerYear
	} else {
		mi = mi + nmonths
	}
	//最后的月份是负数，年份需要减一年月份加12就得到正确的月份，举例：2003年-5月 等于2002年7月
	if mi < 0 {
		yi = yi - 1
		mi = mi + MonthsPerYear
	}
	y := strconv.Itoa(yi)
	m := strconv.Itoa((mi))

	if mi < 10 {
		return y + "0" + m
	}
	return y + m
}

//TimestampToFormat 时间戳(秒)转为格式化字符串"2006-01-02 15:04:05"
func TimestampToFormat(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(FormatLayout)
}

//FormatToTimestamp 格式化字符串转为时间戳-秒
func FormatToTimestamp(timestr string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}
	t, err := time.ParseInLocation(FormatLayout, timestr, loc)

	return t.Unix(), err
}
