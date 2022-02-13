package logger

import "fmt"

// Logger 构造函数
type Logger struct {
}

// NewLog 日志对象
func NewLog() Logger {
	return Logger{}
}

func (l Logger) Debug(msg string) {
	fmt.Println(msg)
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l Logger) Error(msg string) {
	fmt.Println(msg)
}
