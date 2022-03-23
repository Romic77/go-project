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

## logagent

1. 启动`project-02/etcd_demo/main.go`，首先需要把kafka主题等相关信息存入etcd

> key:   chenqi/192.168.5.102
>
> value:	[{"path":"d:/tmp/web.log","topic":"web_log"},{"path":"d:/tmp/redis.log","topic":"redis_log"}]

2. 启动`project-02/logagent/main.go`,创建d:/tmp/web.log、d:/tmp/redis.log日志文件

## log_transfer

1. 启动elasticsearch

   双击`E:\Program Files\es-cluster\node-9201\bin\elasticsearch.bat`

2. 启动`project-02/log_transfer/main.go`

3. 启动kibana-7.8.0-windows-x86_64

   双击`D:\迅雷下载\其他\kibana-7.8.0-windows-x86_64\bin\kibana.bat`

4. 打开http://localhost:5601/app/kibana#/home

5. kibana查看索引为web_log的日志
6. 