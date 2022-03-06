package main

import (
	"fmt"
	"go-project/project-02/http_demo/protocol"
	"net"
	"os"
)

func main() {
	//1. 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("dial 127.0.0.1 failed,error: ", err)
		return
	}
	for i := 0; i < 100; i++ {
		//2. 发送数据
		var msg string
		if len(os.Args) < 2 {
			msg = "hello world"
		} else {
			msg = os.Args[1]
		}
		encode, _ := protocol.Encode(msg)
		conn.Write(encode)
	}
	defer conn.Close()
}
