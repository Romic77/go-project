package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int //需要指定通道中元素的类型

var waitGroup sync.WaitGroup

//make固定类型初始化 channel、map、slice 都是引用类型
//new 值类型
//通道的操作 发送值  ch1 <-1   接收值 x:=<-ch1 关闭 close()

func noBufChannel() {
	defer waitGroup.Done()
	//无缓冲区
	b = make(chan int)
	go func() {
		b <- 10
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	//b <- 10 //因为没有设置缓冲区，导致死锁
}

func bufChannel() {
	defer waitGroup.Done()
	//有缓冲区
	b = make(chan int, 10)
	b <- 10
	x := <-b
	fmt.Println("goroutine从通道b中取到了", x)
}

func main() {
	//通道的初始化
	a = make([]int, 10)

	waitGroup.Add(2)
	go noBufChannel()
	go bufChannel()
	waitGroup.Wait()
}
