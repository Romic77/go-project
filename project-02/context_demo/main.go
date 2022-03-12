package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

//context 就是为了解决goroutine的子线程退出的问题；通过传递环境变量来控制
func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	waitGroup.Add(1)
	go f(ctx)
	cancelFunc()
	waitGroup.Wait()

}

func f(ctx context.Context) {
	defer waitGroup.Done()
	go f2(ctx)
FORLOOP:
	for {
		fmt.Println("调用")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func f2(ctx context.Context) {
FORLOOP:
	for {
		fmt.Println("调用")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func f3() {
	d := time.Now().Add(50 * time.Millisecond)
	//设置取消时间
	ctx, cancelFunc := context.WithDeadline(context.Background(), d)
	//手动调用停止子线程
	defer cancelFunc()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("牛逼")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
