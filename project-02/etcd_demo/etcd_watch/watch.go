package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connection etcd error: %v\n", err)
		return
	}
	fmt.Println("connect to etcd success")

	defer client.Close()

	//watch
	watchChan := client.Watch(context.Background(), "chenqi")
	for response := range watchChan {
		for _, event := range response.Events {
			fmt.Printf("type: %v key: %v value:%v \n", event.Type, string(event.Kv.Key), string(event.Kv.Value))
		}
	}

}
