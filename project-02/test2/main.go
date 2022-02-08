package main

type speak interface {
	speak()
}

type cat struct {
}

type dog struct {
}

type person struct {
}

func (c cat) speak() {
	println("喵喵喵~")
}

func (d dog) speak() {
	println("汪汪汪~")
}
func (p person) speak() {
	println("嘤嘤嘤~")
}

func da(s speak) {
	//接收一个参数，传进来什么，我就打什么
	s.speak()
}

//接口是一个类型，接口里面有方法，传递进去的对象需要包含接口的方法
func main() {
	var c1 cat
	var d1 dog
	var p1 person
	da(c1)
	da(d1)
	da(p1)
}
