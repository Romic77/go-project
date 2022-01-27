package main

func main() {
	add(1, 2)
}

func add(num1 int, num2 int) {
	//go语言遇到defer关键字 语句不会立即执行。会将语句压入栈中，然后继续执行函数后面的语句； 栈的特点是先入后出。 所以先sum、后num2、最后num1
	defer println("num1", num1)
	defer println("num2", num2)
	num1 += 30
	num2 += 30

	sum := num1 + num2
	println(sum)
}
