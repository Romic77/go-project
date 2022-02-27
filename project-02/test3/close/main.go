package main

func main() {
	ch1 := make(chan int, 100)
	/*ch1 <- 1
	ch1 <- 2
	//通道管理也还能读取
	close(ch1)
	for value := range ch1 {
		fmt.Println(value)
	}*/

	go f1(ch1)
}

//单向通道
func f1(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
}
