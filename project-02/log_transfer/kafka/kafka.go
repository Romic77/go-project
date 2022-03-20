package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log_transfer/es"
)

// Init
// @Description: 初始化kafka连接，消费队列数据
// @param address string
// @param topic string
func Init(address string, topic string) (err error) {
	consumer, err := sarama.NewConsumer([]string{address}, nil)
	if err != nil {
		fmt.Printf("failed to start consumer , err:%v\n", err)
		return err
	}
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("failed to get list of partitions , err:%v\n", err)
		return err
	}

	for partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to consume partition %d,err:%v\n", partition, err)
			return err
		}

		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition: %d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				//直接发送给es
				logData := &es.LogData{Topic: topic, Data: string(msg.Value)}

				json.Unmarshal(msg.Value, logData)

				es.SendToESChan(logData)
			}
		}(pc)
	}
	return
}
