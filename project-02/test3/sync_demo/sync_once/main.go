package main

import "sync"

//sync.Once 并行的时候只加载一次的函数，比如close方法等
var waigGroup sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer waigGroup.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer waigGroup.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	f := func() {
		close(ch2)
	}
	once.Do(f)
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	waigGroup.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	waigGroup.Wait()
}
