package main

import (
	"fmt"
	"time"
)

// goroutine
func hello(i int) {
	fmt.Println("hello", i)
}
func main() {
	for i := 0; i < 100; i++ {
		go hello(i)
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			hello(i)
		}(i)
	}

	fmt.Println("main")
	time.Sleep(time.Second * 1)
}
