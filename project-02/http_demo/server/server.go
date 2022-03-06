package main

import (
	"bufio"
	"fmt"
	"go-project/project-02/http_demo/protocol"
	"io"
	"net"
)

func main() {
	//本地端口服务启动
	listen, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:2000 failed, err:", err)
		return
	}
	for {
		//等待别人来跟我建立连接
		conn, err := listen.Accept()
		if err != nil {
			return
		}
		if err != nil {
			fmt.Println("accept failed,err:", err)
			return
		}
		go doRequest(conn)
	}
}

func doRequest(conn net.Conn) {
	//与客户端通信
	reader := bufio.NewReader(conn)
	for {
		decode, err := protocol.Decode(reader)
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("read from conn failed,err:", err)
			return
		}

		fmt.Println(decode)
	}
}
