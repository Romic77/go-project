package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log_transfer/config"
	"log_transfer/es"
	"log_transfer/kafka"
)

var (
	//加载配置文件，初始化LogTransfer
	//new 返回的是对应类型的指针
	cfg = new(config.LogTransfer)
)

func main() {
	//加载配置文件
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Printf("load config err:%v\n", err)
		return
	}
	fmt.Printf("load config success,config=%v\n", cfg)

	//初始化es
	err = es.Init(cfg.ESConfig.Address, cfg.ESConfig.ChanSize, cfg.ESConfig.ConsumerGoroutineNums)
	if err != nil {
		fmt.Printf("init es err:%v\n", err)
		return
	}
	fmt.Println("es init success")

	//初始化kafka
	err = kafka.Init(cfg.KafkaConfig.Address, cfg.KafkaConfig.Topic)
	if err != nil {
		fmt.Printf("init kafka consumer err:%v\n", err)
		return
	}
	fmt.Println("kafka init success")
	select {}
}
