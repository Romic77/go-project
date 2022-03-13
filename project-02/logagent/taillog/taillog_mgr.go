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

	for key, logEntry := range logEntryConf {
		fmt.Printf("key: %v logEntry: %v\n", key, logEntry.Path)

		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		logMgr.taskMap[mk] = tailObj
	}
	go logMgr.run()
}

// run
// @Description:
// @receiver t *tailLogMgr
func (t *tailLogMgr) run() {
	for {
		select {
		case newConfigChan := <-t.newConfigChan:
			for _, conf := range newConfigChan {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.taskMap[mk]
				if ok {
					continue
				} else {
					//1. etcd新增配置
					task := NewTailTask(conf.Path, conf.Topic)
					t.taskMap[mk] = task
				}
			}

			for _, c1 := range t.logEntry {
				isDelete := true
				for _, c2 := range newConfigChan {
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					//停止taillog
					t.taskMap[mk].cancelFunc()
				}
			}
			//2. etcd删除配置
		default:
			time.Sleep(time.Second)
		}
	}
}

func NewConfigChan() chan<- []*etcd.LogEntry {
	return logMgr.newConfigChan
}
