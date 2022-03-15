package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("failed to start consumer , err:%v\n", err)
		return
	}
	partitions, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("failed to get list of partitions , err:%v\n", err)
		return
	}
	fmt.Println(partitions)

	for partition := range partitions {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to consume partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition: %d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}

}
