package config

// 映射的核心是通过后面`ini:`
type AppConfig struct {
	KafkaConfig `ini:"kafka"` //对应的是配置文件的[kafka]
	EtcdConfig  `ini:"etcd"`
}

type KafkaConfig struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

type EtcdConfig struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"key"`
}

type TaillogConfig struct {
	FileName string `ini:"fileName"`
}
