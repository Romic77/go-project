package rpc_design

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type MyInterface interface {
	HelloWorld(name string, resp *string) error
}

func RegisterService(i MyInterface) {
	err := rpc.RegisterName("rpcServer", i)
	if err != nil {
		fmt.Println("注册rpc服务失败!", err)
		return
	}
}

type MyClient struct {
	c *rpc.Client
}

func InitClient(addr string) MyClient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return MyClient{c: conn}
}

//
// HelloWorld
// @Description 实现了类的接口方法
// @receiver m *MyClient
// @param name string
// @param resp *string
// @return error
//
func (m *MyClient) HelloWorld(name string, resp *string) error {
	return m.c.Call("rpcServer.HelloWorld", name, resp)
}
