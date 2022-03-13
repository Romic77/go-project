package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	client *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func Init(addr string, time time.Duration) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: time,
	})
	if err != nil {
		fmt.Printf("connection etcd error: %v\n", err)
		return
	}
	return
}

// 从etcd中通过key获取配置项
func GetByKey(key string) (logEntry []*LogEntry, err error) {
	resp, err := client.Get(context.Background(), key)
	if err != nil {
		fmt.Printf("get to etcd error: %v\n", err)
		return
	}

	for _, value := range resp.Kvs {
		err = json.Unmarshal(value.Value, &logEntry)
		if err != nil {
			fmt.Printf("Unmarshal etcd value failed,error:%v", err)
			return
		}
	}
	return
}

func WatchByKey(key string, newConfigChan chan<- []*LogEntry) {
	watchChan := client.Watch(context.Background(), key)
	for response := range watchChan {
		for _, event := range response.Events {
			fmt.Printf("type: %v key: %v value:%v \n", event.Type, string(event.Kv.Key), string(event.Kv.Value))
			var newConfig []*LogEntry
			//通知别人
			if event.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(event.Kv.Value, &newConfig)
				if err != nil {
					fmt.Printf("unmarshal failed,error:%v\n", err)
					continue
				}
			}

			fmt.Printf("get new conf :%v\n", newConfigChan)
			newConfigChan <- newConfig

		}
	}
}
