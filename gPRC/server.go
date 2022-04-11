package main

import (
	"context"
	"fmt"
	"gRPC/proto"
	"google.golang.org/grpc"
	"net"
)

type UserInfoService struct {
	proto.UnimplementedUserInfoServiceServer
}

func (s *UserInfoService) GetUserInfo(ctx context.Context, req *proto.UserRequest) (resp *proto.UserResponse, err error) {
	name := req.Name
	if name == "zhangsan" {
		resp = &proto.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"sing", "run"},
		}
	}
	return
}

func main() {
	addr := "127.0.0.1:8080"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	server := grpc.NewServer()

	proto.RegisterUserInfoServiceServer(server, &UserInfoService{})

	fmt.Println("服务端启动，地址为：", addr)
	server.Serve(listen)
}
