package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := CopyFile("copy.txt", "src.txt")
	if err != nil {
		fmt.Println("copy file failed,err:", err)
		return
	}

	fmt.Println("copy done")

}

//文件拷贝
func CopyFile(dstName string, srcName string) (written int64, err error) {
	//以读的方式打开源文件
	reader, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed,err:%v", srcName, err)
		return 0, err
	}

	defer reader.Close()

	//写文件
	writer, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open %s failed,err:%v", dstName, err)
		return 0, err
	}
	defer writer.Close()

	return io.Copy(writer, reader)
}
