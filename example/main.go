package main

import (
	"fmt"
	"tools/tools"
)

func main() {

	/************time转换小函数******************/
	//获取当前时间戳-秒单位
	fmt.Printf("now GetNowUnix: %d, \n", tools.GetNowUnix())

	//获取当前时间戳-毫秒单位
	fmt.Printf("now GetNowUnixMilli: %d, \n", tools.GetNowUnixMilli())

	//获取当前时间戳-纳秒单位
	fmt.Printf("now GetNowUnixNano: %d, \n", tools.GetNowUnixNano())

	//获取当前时间字符串格式
	fmt.Printf("now farmat: %s, \n", tools.GetNowFarmat())

	//获取今天日期
	fmt.Printf("today date: %s, \n", tools.GetDateAfterNDays(0))

	//获取三天后的日期
	fmt.Printf("%d days later date: %s, \n", 3, tools.GetDateAfterNDays(3))

	//获取15天前的日期
	fmt.Printf("%d days before date: %s, \n", 15, tools.GetDateAfterNDays(-15))

	//获取当前月份
	fmt.Printf("cur month: %s, \n", tools.GetMonthAfterNMonths(0))

	//获取8个月后的月份
	fmt.Printf("%d months later month: %s, \n", 8, tools.GetMonthAfterNMonths(8))

	//获取197个月前的月份
	fmt.Printf("%d months before month: %s, \n", 197, tools.GetMonthAfterNMonths(-197))

	//时间戳秒 转格式化日期
	var timestamp int64 = 1545469049
	fmt.Printf("timestamp %d --> %s \n", timestamp, tools.TimestampToFormat(timestamp))

	//转格式化日期 转时间戳秒
	timestr := "2012-12-14 13:14:00"
	ti, err := tools.FormatToTimestamp(timestr)
	fmt.Printf("%s -->timestamp %d, err:%v \n", timestr, ti, err)

}
