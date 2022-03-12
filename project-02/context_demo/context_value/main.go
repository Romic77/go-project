package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var waitgroup sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker , trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done")
	waitgroup.Done()
}

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "123123123123")
	waitgroup.Add(1)
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancelFunc()
	waitgroup.Wait()
	fmt.Println("over")
}
