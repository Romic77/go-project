package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU())

	waitGroup.Add(2)
	go a()
	go b()
	fmt.Println(runtime.NumGoroutine())
	waitGroup.Wait()
	fmt.Println("end")
}

func a() {
	defer waitGroup.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer waitGroup.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("b:%d\n", i)
	}
}
