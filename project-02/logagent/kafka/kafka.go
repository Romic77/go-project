package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var (
	//全局kafka连接
	client      sarama.SyncProducer
	logDataChan chan *logData
)

type logData struct {
	topic string
	data  string
}

func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed ,error:", err)
		return
	}

	//初始化logDataChan通道
	logDataChan = make(chan *logData, maxSize)

	//开启kafka通道接收数据，这样使用goroutine 速度快
	go sendToKafka()
	return
}

// sendToKafka 接受通道的消息，然后
func sendToKafka() {
	for {
		select {
		case data := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = data.topic
			msg.Value = sarama.StringEncoder(data.data)

			//发送消息
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed,error:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v err:%v\n", pid, offset, err)
			fmt.Println("发送成功")
		default:
			time.Sleep(time.Second)
		}
	}
}

// SendToChan 消息发送到通道
func SendToChan(topic string, data string) {
	msg := &logData{topic: topic, data: data}
	logDataChan <- msg
}
