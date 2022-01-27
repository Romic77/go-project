package main

func main() {
	//定义匿名函数,并且传入值10，20. 使用result int 接收返回值
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)

	println(result)

	//将匿名函数 复制给sub变量
	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}

	println(sub(1, 2))

	//调用全局匿名函数
	println(Result(2, 2))

}

// Result 定义全局匿名函数，然后复制给Result，全局变量不能使用:=
var Result = func(num1 int, num2 int) int {
	return num1 - num2
}
