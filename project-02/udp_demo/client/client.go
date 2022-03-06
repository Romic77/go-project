package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		fmt.Println("连接服务器失败，err:", err)
		return
	}
	defer socket.Close()

	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入内容")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))

		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("redv reply msg failed ,err", err)
			return
		}
		fmt.Println("收到回复消息:", string(reply[:n]))
	}
}
