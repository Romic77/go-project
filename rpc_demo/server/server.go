package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type World struct {
}

func (w *World) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好!"
	return nil
}

//
// main
// @Description 这个注释还是有点看不清
//
func main() {
	//err := rpc.RegisterName("rpcServer", new(World))
	err := rpc.RegisterName("rpcServer", new(World))
	if err != nil {
		fmt.Println("注册rpc服务失败!", err)
		return
	}

	listen, _ := net.Listen("tcp", "127.0.0.1:8080")

	conn, _ := listen.Accept()

	//rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)
	defer listen.Close()

}
