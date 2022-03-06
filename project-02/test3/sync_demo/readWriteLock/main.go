package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x         = 0
	rwLock    sync.RWMutex
	waitGroup sync.WaitGroup
)

func read() {
	defer waitGroup.Done()
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond * 5)
	rwLock.RUnlock()
}

func write() {
	defer waitGroup.Done()
	rwLock.Lock()
	x++
	rwLock.Unlock()
}
func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go write()

	}

	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go read()
	}
	fmt.Println(time.Now().Sub(start))
	waitGroup.Wait()
}
