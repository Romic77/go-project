package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

//
// main
// @Description 注释亮一点
//
func main() {
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Dial error", err)
		return
	}
	defer conn.Close()

	var reply string

	// 入参第一个传递 需要调用的 服务名.函数名
	// 服务名就是 rpc.RegisterName第一个入参 为 hello
	err = conn.Call("rpcServer.HelloWorld", "李白", &reply)
	if err != nil {
		fmt.Println("call error", err)
		return
	}

	fmt.Println(reply)

}
