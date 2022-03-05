package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
使用goroutine和channel实现一个计算int64随机数各位数和的程序
1. 开启一个goroutine循环生成int64随机数，发送到jobChan
2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3. 主goroutine从resultChan取出结果并打印
*/

type job struct {
	num int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var waitGroup sync.WaitGroup

// 往jobChan写入数据
func producer(jobChan chan<- *job) {
	defer waitGroup.Done()
	//循环生成int64类型的随机数，发送到jobChan
	rand.Seed(time.Now().UnixNano())
	for {
		int63 := rand.Int63()
		/*//通过new来进行初始化结构体
		newJob := new(job)
		newJob.num = int63*/

		//通过&job来进行初始化和赋值
		newJob := &job{num: int63}
		jobChan <- newJob
	}
}

// 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
func consumer(jobChan <-chan *job, resultChan chan<- *result) {
	for {
		//把通道里面值取出来，通道里面值是job的struct，所以用job接收
		job := <-jobChan
		sum := int64(0)
		n := job.num
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	waitGroup.Add(1)
	go producer(jobChan)
	//开启24个goroutine执行消费者
	for i := 0; i < 24; i++ {
		go consumer(jobChan, resultChan)
	}
	// 主goroutine从resultChan取出结果并打印
	for value := range resultChan {
		fmt.Printf("value:%d sum:%d\n", value.job.num, value.sum)
	}

	waitGroup.Wait()
}
