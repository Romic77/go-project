package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type MyHandler struct {
	Title string
}

// 接受消息

func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s received from %v,message: %v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return
}

//初始化消费者
func initConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed: %v\n", err)
		return err
	}
	consumerObj := &MyHandler{
		Title: "深圳福田",
	}
	//这里才是设计的核心，因为HandlerMessage的作用域是MyHandler的方法
	consumer.AddHandler(consumerObj)
	err = consumer.ConnectToNSQLookupd(address)
	if err != nil {
		return err
	}
	return
}

func main() {
	err := initConsumer("topic_demo", "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Printf("init consumer failed: %v\n", err)
		return
	}
	//make slice/map/chan
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP)
	<-c
}
