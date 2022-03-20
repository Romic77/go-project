# Elk日志收集项目

## Kafka

1. 下载kafka安装包

   `kafka_2.12-3.1.0`

2. cmd启动zookeeper

   `.\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties `

3. 新启cmd，启动kafka-server

   `.\bin\windows\kafka-server-start.bat .\config\server.properties`

## Etcd

1. 下载etcd-v3.5.0-windows-amd64

2. 启动etcd

   双击`etcd.exe`

## Logagent

1. 启动`project-02/etcd_demo/main.go`，首先需要把kafka主题等相关信息存入etcd

> key:   chenqi/192.168.5.102
>
> value:	[{"path":"d:/tmp/web.log","topic":"web_log"},{"path":"d:/tmp/redis.log","topic":"redis_log"}]

2. 启动`project-02/logagent/main.go`,创建d:/tmp/web.log、d:/tmp/redis.log日志文件