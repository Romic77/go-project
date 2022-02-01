package main

import "fmt"

func main() {
	var a = []int{1, 4, 7, 2, 5, 8}
	//定义切片b,类型为int，长度为10
	var b = make([]int, 10)
	//把a里面的元素给b
	copy(b, a)

	fmt.Println(b)
}
