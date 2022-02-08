package main

import "fmt"

//接口，跑的方法
type car interface {
	run()
}

//结构体类似java的类， 类实现car的接口
type ferrari struct {
	brand string
}

type lamborghini struct {
	brand string
}

type bicycle struct {
	brand string
}

func (f ferrari) run() {
	fmt.Printf("%s开启\n", f.brand)
}

func (l lamborghini) run() {
	fmt.Printf("%s开启\n", l.brand)
}

func (b bicycle) run() {
	fmt.Printf("%s开启\n", b.brand)
}

func startRun(c car) {
	c.run()
}

func main() {
	f := ferrari{brand: "法拉利"}
	l := lamborghini{brand: "兰博基尼"}
	b := bicycle{brand: "自行车"}

	startRun(f)
	startRun(l)
	startRun(b)
}
