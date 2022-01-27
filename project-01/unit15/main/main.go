package main

//函数名是getSum(), 返回值是 func(int) int
func getSum() func(int) int {
	//闭包中使用的变量会一直保存在内存中，意思就是sum一直会在内存中，所以f1=1 f2=1+2=3 f3=3+3=6
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}
}

var sum int = 0

func getSum1(num int) int {
	return num + sum
}

func main() {
	f := getSum()
	//结果是1
	println(f(1))
	//正常结果是2，执行出来是3
	println(f(2))
	println(f(3))

	//不使用闭包则需要使用全局变量等完成
	sum = getSum1(1)
	sum = getSum1(2)
	sum = getSum1(3)
	println(sum)

}
