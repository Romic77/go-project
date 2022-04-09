package main

import (
	"context"
	"fmt"
	"gRPC/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	client := proto.NewUserInfoServiceClient(conn)
	req := new(proto.UserRequest)
	req.Name = "zhangsan"
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		return
	}

	fmt.Printf("响应结果:%v\n", response)
}
