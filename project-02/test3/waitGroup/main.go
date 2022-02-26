package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func f() {
	//不加种子 那么随机数每次都是一样的，需要使用seed方法
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}
}

func f1(i int) {
	//类似countDownLatch.countDown()
	defer waitGroup.Done()
	fmt.Println(i)

}

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		f1(i)
	}
	waitGroup.Wait()
	fmt.Println("main end")
}
