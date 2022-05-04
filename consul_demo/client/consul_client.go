package main

import (
	"consul_demo/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func main() {

	consulConfig := api.DefaultConfig()
	consulClient, _ := api.NewClient(consulConfig)

	serviceEntries, _, err := consulClient.Health().Service("consulService", "", true, nil)
	if err != nil {
		fmt.Println("consulService服务不可用")
		return
	}
	addr := serviceEntries[0].Service.Address + ":" + strconv.Itoa(serviceEntries[0].Service.Port)
	grpcConn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("打开grpc连接")
		return
	}

	grpcClient := pb.NewHelloClient(grpcConn)

	result, _ := grpcClient.SayHello(context.TODO(), &pb.Person{Name: "陈奇consul", Age: 30})

	fmt.Println(result)
	defer grpcConn.Close()
}
