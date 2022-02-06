package main

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	addr address
}

type company struct {
	name    string
	address //匿名嵌套结构体 ,个人不建议使用匿名嵌套。
}

func main() {
	p1 := person{name: "zhangsan", age: 12, addr: address{
		province: "广东",
		city:     "深圳",
	}}

	p2 := company{name: "wezhuiyiu", address: address{
		province: "广东",
		city:     "深圳",
	}}

	println(p1.name, p1.addr.city)
	println(p2.name, p2.city)
}
