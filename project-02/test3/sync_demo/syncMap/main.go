package main

import (
	"fmt"
	"strconv"
	"sync"
)

var syncMap = sync.Map{}
var waitGroup sync.WaitGroup

func main() {
	for i := 0; i < 200; i++ {
		waitGroup.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			//store存储值
			syncMap.Store(key, n)
			//load取值返回value和bool 是否取值成功
			value, _ := syncMap.Load(key)
			fmt.Println("k=:%v,v:=%v\n", key, value)
			defer waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
}
