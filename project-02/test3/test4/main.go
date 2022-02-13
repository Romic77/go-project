package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//时间间隔
	fmt.Println(now.Second())
	//now+24小时
	fmt.Println(now.Add(time.Hour * 24))

	//定时器
	/*timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t)
	}*/

	//格式化时间 把语言中时间对象 转换成字符串类型的时间
	fmt.Println(now.Format("2006-01-02"))
	//固定格式 2006 01 02 03 04 05
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	//将字符串解析为时间
	fmt.Println(time.Parse("2006-01-02 15:04:05", "2022-02-13 14:23:00"))

	//当前时间和2020-07-03 相差多少时间
	nextYear, err := time.Parse("2006-01-02", "2020-07-03")
	if err != nil {
		return
	}

	d := now.Sub(nextYear)
	fmt.Println(d)
	//
	n := 5
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("5秒钟过去了")

	f1()
}

func f1() {

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}

	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-02-13 16:53:00", location)
	if err != nil {
		return
	}
	fmt.Println(timeObj)

}
