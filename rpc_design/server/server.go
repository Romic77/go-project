package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"rpc_design"
)

type World struct {
}

func (w *World) HelloWorld(name string, resp *string) error {
	*resp = "你好，" + name
	return nil
}

func main() {
	rpc_design.RegisterService(new(World))

	listen, _ := net.Listen("tcp", "127.0.0.1:8080")

	conn, _ := listen.Accept()
	var tmp [512]byte
	conn.Read(tmp[0:])
	fmt.Println(string(tmp[0:]))

	//rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)

	defer listen.Close()

}
