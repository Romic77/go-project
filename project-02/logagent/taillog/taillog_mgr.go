package taillog

import (
	"fmt"
	"logagent/etcd"
	"time"
)

var (
	logMgr *tailLogMgr
)

//  tailLogMgr
//  @Description: 日志管理者
type tailLogMgr struct {
	logEntry      []*etcd.LogEntry
	taskMap       map[string]*TailTask
	newConfigChan chan []*etcd.LogEntry
}

// Init
// @Description: 初始化
// @param logEntryConf []*etcd.LogEntry
func Init(logEntryConf []*etcd.LogEntry) {
	logMgr = &tailLogMgr{
		logEntry:      logEntryConf,
		taskMap:       make(map[string]*TailTask, 16),
		newConfigChan: make(chan []*etcd.LogEntry), //无缓冲区的通道
	}

	for key, value := range logEntryConf {
		fmt.Printf("key: %v value: %v\n", key, value.Path)

		NewTailTask(value.Path, value.Topic)
	}
}

// run
// @Description:
// @receiver t *tailLogMgr
func (t *tailLogMgr) run() {
	for {
		select {
		case newConfigChan := <-t.newConfigChan:
			fmt.Println(newConfigChan)
		//
		default:
			time.Sleep(time.Second)
		}
	}
}

func NewConfigChan() chan<- []*etcd.LogEntry {
	return logMgr.newConfigChan
}
