package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	//获取当前时间的时分秒
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	dateStr := fmt.Sprintf("当前年月日：%d %d %d 时分秒：%d %d %d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	println(dateStr)

	//2006/01/02 15/04/05 固定格式化字符串，不能改变。庆祝go出生点的日期
	dateStr2 := now.Format("2006/01/02 15/04/05")
	println(dateStr2)

	println(now.Format("2006-01-02 15-04-05"))
}
