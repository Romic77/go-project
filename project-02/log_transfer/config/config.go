package config

// 映射的核心是通过后面`ini:`
type LogTransfer struct {
	KafkaConfig `ini:"kafka"` //对应的是配置文件的[kafka]
	ESConfig    `ini:"es"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESConfig struct {
	Address               string `ini:"address"`
	ChanSize              int    `ini:"chan_size"`
	ConsumerGoroutineNums int    `ini:"consumer_goroutine_nums"`
}
