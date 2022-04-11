package main

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"sync"
)

func main() {
	c := &config.AppConfig{
		AppID:          "rpa-control",
		Cluster:        "default",
		IP:             "http://172.16.30.73:8080",
		NamespaceName:  "sessionmanager.json",
		IsBackupConfig: true,
		//Secret:         "6ce3ff7e96a24335a9634fe9abca6d51",
	}

	//agollo.SetLogger(&DefaultLogger{})
	c2 := &CustomChangeListener{}
	c2.wg.Add(5)

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	cache := client.GetConfigCache(c.NamespaceName)
	value, _ := cache.Get("sessionmanager")
	fmt.Println(value)

	client.AddChangeListener(c2)

	c2.wg.Wait()
	writeConfig(c.NamespaceName, client)
}

func writeConfig(namespace string, client agollo.Client) {
	cache := client.GetConfigCache(namespace)
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		return true
	})
}

type CustomChangeListener struct {
	wg sync.WaitGroup
}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	//write your code here
	fmt.Println(changeEvent.Changes)
	for key, value := range changeEvent.Changes {
		fmt.Println("change key : ", key, ", value :", value)
	}
	fmt.Println(changeEvent.Namespace)
	c.wg.Done()
}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	//write your code here
	/*for {
		//Use your apollo key to test
		cache := client.GetConfigCache(c.NamespaceName)
		value, _ := cache.Get("sessionmanager.conf")
		fmt.Println(value)
		time.Sleep(time.Second * 5)
	}*/
	fmt.Println("onNewestChange:", event.Changes)
	c.wg.Done()
}
