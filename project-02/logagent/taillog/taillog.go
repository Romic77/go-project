package taillog

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"logagent/kafka"
	"time"
)

var (
	tailObj *tail.Tail
	LogChan chan string
)

type TailTask struct {
	path       string
	topic      string
	instance   *tail.Tail
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path string, topic string) (tailObj *TailTask) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancelFunc,
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
		case <-t.ctx.Done():
			fmt.Printf("tail task:%s %s 结束了\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines:
			//2.发送到kafka通道
			//kafka.SentToKafka(t.topic, line.Text)
			kafka.SendToChan(t.topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
