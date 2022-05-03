package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"rpc_design/pb"
)

func main() {
	grpcConn, _ := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())

	defer grpcConn.Close()
	grpcClient := pb.NewSayNameClient(grpcConn)

	t, err := grpcClient.SayHello(context.TODO(), &pb.Teacher{Name: "mic", Age: 35})
	fmt.Println(t, err)
}
