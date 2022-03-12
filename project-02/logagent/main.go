package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logagent/config"
	"logagent/kafka"
	"logagent/taillog"
	"time"
)

var (
	cfg = new(config.AppConfig)
)

func main() {
	initConfig()

	//1. 初始化一个kafka链接
	err := kafka.Init([]string{cfg.KafkaConfig.Address})
	if err != nil {
		fmt.Printf("init kafka failed: %v\n", err)
		return
	}
	fmt.Println("init kafka success")
	err = taillog.Init(cfg.TaillogConfig.FileName)
	if err != nil {
		fmt.Printf("Init taillog failed: %v\n", err)
		return
	}
	run()

}

//加载配置文件
func initConfig() {
	//加载配置文件
	err := ini.MapTo(cfg, "./config/config.ini")

	if err != nil {
		fmt.Printf("init load config error: %v", err)
		return
	}
}

func run() {
	//1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2.发送到kafka
			kafka.SentToKafka(cfg.KafkaConfig.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}
