package main

import (
	"go-project/project-02/test3/test5/logger"
	"log"
	"os"
)

func main() {
	logger := logger.NewLog("info")
	logger.Debug("这是一条Debug日志")
	logger.Info("这是一条Info日志")
	logger.Error("这是一条Error日志")
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
