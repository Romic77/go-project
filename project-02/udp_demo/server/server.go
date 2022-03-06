package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		fmt.Println("listent udp failed,err:", err)
		return
	}
	defer conn.Close()
	//不需要建立连接，直接发送数据
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from udp failed,err:", err)
			return
		}
		reply := strings.ToUpper(string(data[:n]))
		fmt.Println(reply)
		//发送数据
		conn.WriteToUDP([]byte(reply), addr)
	}

}
