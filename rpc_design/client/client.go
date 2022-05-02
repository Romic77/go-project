package main

import (
	"fmt"
	"go-project/rpc_design"
)

func main() {
	myClient := rpc_design.InitClient("127.0.0.1:8080")

	var resp string
	err := myClient.HelloWorld("杜甫1", &resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
