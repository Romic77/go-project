package main

import (
	"fmt"
	"go-project/project-02/test3/test5/myLogger"
	"log"
	"os"
	"time"
)

func main() {
	//printConsoleLog()
	printFileLog()
}

//测试往console打印日志
func test() {
	writter, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	log.SetOutput(writter)

	for {
		log.Println("这是一条日志")
	}
}

func printConsoleLog() {
	for {
		logger := myLogger.NewLog("info")
		logger.Debug("这是一条Debug日志")
		logger.Info("这是一条Info日志")
		logger.Error("这是一条Error日志")
		//每5秒打印日志
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func printFileLog() {
	fileLogger := myLogger.NewFileLogger("info", "./", "cq.log", 10*1024)
	for {
		fileLogger.Debug("这是一条Debug日志")
		fileLogger.Info("这是一条Info日志")
		fileLogger.Error("这是一条Error日志")
		//每5秒打印日志
		//time.Sleep(time.Duration(2) * time.Second)
	}
}

//日志切割案例
func logSplit() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
	}
	// 1.文件对象的类型
	fmt.Printf("%T\n", fileObj)
	// 2.获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
	}
	fmt.Printf("文件大小是：%d\n", fileInfo.Size())
}
