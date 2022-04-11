package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"time"
)

func main() {
	c := &config.AppConfig{
		AppID:          "rpa-control",
		Cluster:        "default",
		IP:             "http://172.16.30.73:8080",
		NamespaceName:  "Test1.springboot",
		IsBackupConfig: true,
		//Secret:         "6ce3ff7e96a24335a9634fe9abca6d51",
	}

	//agollo.SetLogger(&DefaultLogger{})

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	fmt.Println("初始化Apollo配置成功")

	for {
		//Use your apollo key to test
		cache := client.GetConfigCache(c.NamespaceName)
		value, _ := cache.Get("public_key")
		fmt.Println(value)
		time.Sleep(time.Second * 5)
	}
}
