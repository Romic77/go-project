package main

func calc(num1 int, num2 int) int {

	return num1 + num2
}

func calc1(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func main() {
	println(calc(1, 2))
	println(calc1(1, 2))
}
