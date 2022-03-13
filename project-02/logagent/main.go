package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logagent/config"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"sync"
	"time"
)

var (
	cfg       = new(config.AppConfig)
	waitgroup sync.WaitGroup
)

func main() {
	//加载配置文件
	initConfig()

	//1. 初始化一个kafka链接
	err := kafka.Init([]string{cfg.KafkaConfig.Address}, cfg.KafkaConfig.ChanMaxSize)
	if err != nil {
		fmt.Printf("init kafka failed: %v\n", err)
		return
	}
	fmt.Println("init kafka success")

	err = etcd.Init(cfg.EtcdConfig.Address, time.Duration(cfg.EtcdConfig.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("connection etcd error: %v\n", err)
		return
	}
	fmt.Println("connect to etcd success")

	//从etcd中获取日志收集的配置信息
	logEntries, err := etcd.GetByKey(cfg.EtcdConfig.Key)
	if err != nil {
		fmt.Printf("etcd getByKey failed: %v\n", err)
		return
	}
	fmt.Printf("etcd getByKey success,%v\n", logEntries)

	taillog.Init(logEntries)

	newConfigChan := taillog.NewConfigChan()
	waitgroup.Add(1)
	go etcd.WatchByKey(cfg.EtcdConfig.Key, newConfigChan)
	waitgroup.Wait()

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
