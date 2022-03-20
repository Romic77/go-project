package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

var (
	client *elastic.Client
	//传递LogData的指针
	ch chan *LogData
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

// Init
// @Description: 初始化es客户端
// @param address string
// @return err error
func Init(address string, chanSize int, consumerGoroutineNums int) (err error) {
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		fmt.Println("es connection error: ", err)
	}
	ch = make(chan *LogData, chanSize)
	for i := 0; i < consumerGoroutineNums; i++ {
		go SentToEs()
	}
	return
}

// SendToESChan
// @Description: 发送消息到管道
// @param logData *LogData
func SendToESChan(logData *LogData) {
	ch <- logData
}

// SentToEs
// @Description:从通道里面取消息
func SentToEs() {
	for {
		select {
		case logData := <-ch:
			put, err := client.Index().Index(logData.Topic).BodyJson(logData).Do(context.Background())
			if err != nil {
				fmt.Println("SendToES error: ", err)
			}
			fmt.Printf("Index user %s to index %s,type %s\n", put.Id, put.Index, put.Type)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
