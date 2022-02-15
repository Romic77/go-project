package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
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

// 获取日志调用主机、文件名、行号
func getInfo(skip int) (funcName string, fileName string, lineNo int) {
	pc, fileName, line, ok := runtime.Caller(skip)

	if !ok {
		fmt.Printf("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	return funcName, path.Base(fileName), line
}
