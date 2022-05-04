package main

import (
	"consul_demo/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
)

type Children struct {
}

//
// SayHello
// @Description 绑定类方法，实现SayHello接口
// @receiver c *Children
// @param ctx context.Context
// @param p *pb.Person
// @return *pb.Person
// @return error
//
func (c *Children) SayHello(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = "hello " + p.GetName()
	return p, nil
}

func main() {
	//初始化consul配置
	consulConfig := api.DefaultConfig()
	//获取consul操作对象
	client, _ := api.NewClient(consulConfig)

	registerService := api.AgentServiceRegistration{ID: "consul_demo", Tags: []string{"consul"}, Name: "consulService", Port: 8081, Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			TCP:      "127.0.0.1:8081",
			Timeout:  "5s",
			Interval: "5s"},
	}
	client.Agent().ServiceRegister(&registerService)

	//生成grpcServer
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))

	//设置监听，指定ip和端口
	listen, _ := net.Listen("tcp", "127.0.0.1:8081")

	defer listen.Close()
	fmt.Println("服务启动成功")
	grpcServer.Serve(listen)

}
