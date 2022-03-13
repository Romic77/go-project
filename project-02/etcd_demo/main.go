package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	putWeb()
	//putRedis()
	//DeleteRedis()
	//client.Delete(context.Background(), "chenqi")
}
func putWeb() {
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
	//put insert+update
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"d:/tmp/web.log","topic":"web_log"}]`
	_, err = client.Put(ctx, "chenqi", value)
	cancelFunc()
	if err != nil {
		fmt.Printf("put to etcd error: %v\n", err)
		return
	}

	//get
	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Second)
	resp, errs := client.Get(ctx, "chenqi")
	cancelFunc()
	if errs != nil {
		fmt.Printf("get to etcd error: %v\n", err)
		return
	}

	for _, value := range resp.Kvs {
		fmt.Printf("%s:%s\n", value.Key, value.Value)
	}
}

func putRedis() {
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
	//put insert+update
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"d:/tmp/web.log","topic":"web_log"},{"path":"d:/tmp/redis.log","topic":"redis_log"}]`
	_, err = client.Put(ctx, "chenqi", value)
	cancelFunc()
	if err != nil {
		fmt.Printf("put to etcd error: %v\n", err)
		return
	}

	//get
	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Second)
	resp, errs := client.Get(ctx, "chenqi")
	cancelFunc()
	if errs != nil {
		fmt.Printf("get to etcd error: %v\n", err)
		return
	}

	for _, value := range resp.Kvs {
		fmt.Printf("%s:%s\n", value.Key, value.Value)
	}
}

// DeleteRedis
// @Description: 删除redis
func DeleteRedis() {
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

	client.Delete(context.Background(), "chenqi")
}
