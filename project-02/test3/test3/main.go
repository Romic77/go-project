package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//readFromFile1()
	//readFromBuffio()
	//readFromIoutil()
	//writeFile()
	//writeByBufio()
	writeByIoutil()
}

//ioUtil 读文件
func readFromIoutil() {
	//打开文件
	file, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("open filed,err:%v\n", err)
		return
	}
	fmt.Println(string(file))

}

//一行一行读使用bufio
func readFromBuffio() {
	//打开文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open filed,err:%v\n", err)
		return
	}
	//记得关闭文件
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
		}

		if err != nil {
			fmt.Printf("read from file failed,err:%v", err)
			return
		}

		fmt.Print(line)
	}

}

func readFromFile1() {
	//打开文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open filed,err:%v\n", err)
		return
	}
	//记得关闭文件
	defer file.Close()

	//读文件
	var tmp [128]byte
	for {
		//第一个返回值表示读取文件的长度，第二个返回值是异常
		read, err := file.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("文件读完了")
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v", err)
			return
		}

		fmt.Printf("读了%d个字节\n", read)
		fmt.Println(string(tmp[:]))
		//读完了 直接return
		if read < 128 {
			return
		}
	}
}

//写文件
func writeFile() {
	file, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer file.Close()

	file.Write([]byte("你在干什么\n"))
	file.WriteString("牛逼\n")

}

//bufio 写文件
func writeByBufio() {
	file, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		writer.WriteString("hello bro\n")
	}

	writer.Flush()
}

//ioutil.Writefile
func writeByIoutil() {
	err := ioutil.WriteFile("./xx.txt", []byte("呆子"), 0666)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
}
