package main

import (
	"fmt"
	"logagent/kafka"
	"logagent/taillog"
	"time"
)

func main() {
	//1. 初始化一个kafka链接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init kafka failed: %v\n", err)
		return
	}
	fmt.Println("init kafka success")
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("Init taillog failed: %v\n", err)
		return
	}
	run()

}

func run() {
	//1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2.发送到kafka
			kafka.SentToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}
