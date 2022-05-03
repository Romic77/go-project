package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"rpc_design/pb"
)

type Children struct {
}

func (c *Children) SayHello(ctx context.Context, t *pb.Teacher) (*pb.Teacher, error) {
	t.Name += " is sleeping"
	return t, nil
}

func main() {
	//初始化一个grpc对象
	grpcServer := grpc.NewServer()

	pb.RegisterSayNameServer(grpcServer, new(Children))

	listen, _ := net.Listen("tcp", "127.0.0.1:8081")
	defer listen.Close()

	grpcServer.Serve(listen)
}
