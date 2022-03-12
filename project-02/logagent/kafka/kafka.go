package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	//全局kafka连接
	client sarama.SyncProducer
)

func Init(addrs []string) (err error) {
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
	return
}

func SentToKafka(topic, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,error:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v err:%v\n", pid, offset, err)
	fmt.Println("发送成功")
}
