package main

import (
	"fmt"
	"sync"
)

//锁-互斥锁(数据库悲观锁)
var x = 0
var waitgroup sync.WaitGroup

// 悲观锁，类似ReentrantLock
var lock sync.Mutex

func add() {
	defer waitgroup.Done()
	lock.Lock()
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	lock.Unlock()
}

func main() {
	waitgroup.Add(2)
	go add()
	go add()
	waitgroup.Wait()

	fmt.Println(x)
}
