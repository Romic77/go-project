package config

// 映射的核心是通过后面`ini:`
type AppConfig struct {
	KafkaConfig   `ini:"kafka"`   //对应的是配置文件的[kafka]
	TaillogConfig `ini:"taillog"` //对应的是配置文件的[taillog]
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type TaillogConfig struct {
	FileName string `ini:"fileName"`
}
