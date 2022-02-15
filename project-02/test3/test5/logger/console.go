package logger

import (
	"fmt"
	"time"
)

// NewLog 日志对象
func NewLog(levelStr string) Logger {
	logLevel := parseLogLevel(levelStr)
	//结构体初始化
	return Logger{logLevel: logLevel}
}

func (l Logger) enable(level LogLevel) bool {
	//当前等级和方法等级比较
	//debug(1) <= info(2)
	return l.logLevel <= level
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

func log(logLevel LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s \n", now.Format("2006-01-02 15:04:05"), getLogString(logLevel), fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}
