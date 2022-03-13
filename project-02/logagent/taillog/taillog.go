package taillog

import (
	"github.com/hpcloud/tail"
	"logagent/kafka"
	"time"
)

var (
	tailObj *tail.Tail
	LogChan chan string
)

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path string, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.Init()
	return
}

func (t *TailTask) Init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	t.instance, _ = tail.TailFile(t.path, config)
	go t.run()
}

func (t *TailTask) ReadChan() chan *tail.Line {
	return t.instance.Lines
}

func (t *TailTask) run() {
	//1.读取日志
	for {
		select {
		case line := <-t.instance.Lines:
			//2.发送到kafka通道
			//kafka.SentToKafka(t.topic, line.Text)
			kafka.SendToChan(t.topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
