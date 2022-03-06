package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//原子操作
var x int64
var waitgroup sync.WaitGroup

func add() {
	defer waitgroup.Done()
	//原子相加
	atomic.AddInt64(&x, 1)
}

func main() {
	waitgroup.Add(2)
	go add()
	go add()
	waitgroup.Wait()
	fmt.Println(x)
}
