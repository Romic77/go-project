package main

import "fmt"

type animal struct {
	name string
}

//移动的方法
func (a animal) move() {
	fmt.Printf("%s会动！", a.name)
}

type dog struct {
	feet uint8
	animal
}

//方法是作用在特定类型的函数
//接收者表示的是调用该方法的具体类型变量，多用类型的首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s汪汪汪！", d.name)
}

func main() {
	d := dog{feet: 4, animal: animal{
		name: "狗",
	}}
	println(d.name)
	d.wang()
	d.move()
}
