package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启3个goroutine
	for i := 0; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	//5个任务
	for i := 0; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	//输出结果
	for i := 0; i <= 5; i++ {
		<-results
	}
}

// jobs 是只读的通道
// results 是只写的通道
func worker(id int, jobs <-chan int, results chan<- int) {
	for value := range jobs {
		fmt.Printf("worker:%d start job:%d \n", id, value)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, value)
		results <- value * 2
	}
}
