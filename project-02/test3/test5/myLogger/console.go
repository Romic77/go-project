package myLogger

import (
	"fmt"
	"time"
)

// NewLog 日志对象
func NewLog(levelStr string) ConsoleLogger {
	logLevel := ParseLogLevel(levelStr)
	//结构体初始化
	return ConsoleLogger{logLevel: logLevel}
}

func (c ConsoleLogger) enable(level LogLevel) bool {
	//当前等级和方法等级比较
	//debug(1) <= info(2)
	return c.logLevel <= level
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	}
	return "DEBUG"
}

func (c ConsoleLogger) Debug(msg string) {
	c.log(DEBUG, msg)
}

func (c ConsoleLogger) Info(msg string) {
	c.log(INFO, msg)
}

func (c ConsoleLogger) Error(msg string) {
	c.log(ERROR, msg)
}

func (c ConsoleLogger) log(logLevel LogLevel, msg string) {
	if c.enable(logLevel) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s \n", now.Format("2006-01-02 15:04:05"), getLogString(logLevel), fileName, funcName, lineNo, msg)
	}
}
