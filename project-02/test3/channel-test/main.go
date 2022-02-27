package main

import (
	"fmt"
	"sync"
)

//channel 练习
//1. 启动一个goroutine，生成100个数发送到ch1
//2. 启动一个goroutine，从ch1中取值，计算其平方放到ch2中

var waitGroup sync.WaitGroup
var once sync.Once

func f1(ch1 chan int) {
	defer waitGroup.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

//forr遍历map的时候 是key，value。遍历channel的时候只有value
func f2(ch1, ch2 chan int) {
	defer waitGroup.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() {
		close(ch2)
	})
}

func f3(ch1, ch2 chan int) {
	defer waitGroup.Done()
	for value := range ch1 {
		ch2 <- value * value
	}
	once.Do(func() {
		close(ch2)
	})
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	waitGroup.Add(2)
	go f1(a)
	//go f2(a, b)
	//go f2(a, b)
	go f3(a, b)
	waitGroup.Wait()

	for value := range b {
		fmt.Println(value)
	}

}
