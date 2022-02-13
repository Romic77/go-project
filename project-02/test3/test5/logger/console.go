package logger

import (
	"fmt"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	ERROR
)

// Logger 构造函数
type Logger struct {
	logLevel LogLevel
}

func parseLogLevel(logLevel string) LogLevel {
	switch strings.ToLower(logLevel) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "error":
		return ERROR
	default:
		return UNKNOWN
	}
}

// NewLog 日志对象
func NewLog(levelStr string) Logger {
	logLevel := parseLogLevel(levelStr)
	return Logger{
		logLevel: logLevel,
	}
}

func (l Logger) enable(level LogLevel) bool {
	//当前等级和方法等级比较
	//debug(1) <= info(2)
	return l.logLevel <= level
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [Debug] %s \n", now.Format("2006-01-02 15:04:05"), msg)
	}

}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [Info] %s \n", now.Format("2006-01-02 15:04:05"), msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [Error] %s \n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
