package main

import "fmt"

//num1 num2数据交换位置
//为什么没有起作用？因为num1,num2的生命周期只是在exchangeNum函数内，没有返回。所以不会改变main的函数的真实值
func exchangeNum(num1 int, num2 int) {
	var num3 int = num2
	num2 = num1
	num1 = num3
}

func exchangeNum1(num1 int, num2 int) (int, int) {
	return num2, num1
}

func exchangeNum2(num1 *int, num2 *int) {
	//*num2 就是真实值
	/*var num3 int = *num2
	*num2 = *num1
	*num1 = num3*/

	*num1, *num2 = *num2, *num1

}

func main() {
	num1 := 1
	num2 := 2
	/*exchangeNum(num1, num2)
	fmt.Printf("num1=%v,num2=%v", num1, num2)
	num1, num2 = exchangeNum1(num1, num2)
	fmt.Printf("\nnum1=%v,num2=%v", num1, num2)
	*/
	exchangeNum2(&num1, &num2)
	fmt.Println("num1:", num1, ";num2:", num2)

	//test(1)
}

//可变参数 N个int参数
func test(args ...int) {
	//将可变参数当做切片来处理
	for i := 0; i < len(args); i++ {
		println(args[i])
	}
}
