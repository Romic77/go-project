package main

func test(num int) {
	println(num)
}

func test2(num1 int, num2 float32, testFunc func(int)) {
	println("调用了")
}

//自定义类型
type myFunc func(int)

func test3(num1 int, num2 float32, testFunc myFunc) {
	println("调用了自定义函数")
}

func test4(num1 int, num2 int) (sum int, sub int) {
	sub = num1 - num2
	sum = num1 + num2
	return
}

func main() {
	a := test
	//a(10)
	test2(10, 3.19, a)

	//给 int类型起了别名叫myInt类型
	type myInt int

	var num1 myInt = 12
	println(num1)
	var num2 int = 30
	num2 = int(num1)
	println(num2)

	test3(10, 3.19, a)

	println(test4(1, 2))
}
