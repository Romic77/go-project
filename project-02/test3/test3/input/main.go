package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	useScan()
	useBufio()
}

func useScan() {
	var str string
	fmt.Print("请输入内容:")
	fmt.Scanln(&str)
	fmt.Printf("请输入的内容是:%s\n", str)

}

func useBufio() {
	var str string
	fmt.Print("请输入内容:")
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	fmt.Printf("请输入的内容是:%s\n", str)
}
