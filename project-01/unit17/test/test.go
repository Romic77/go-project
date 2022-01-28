package main

import (
	"errors"
	"fmt"
)

func main() {
	//println(test(10, 0))

	result, err := test1(10, 2)
	println(result, err)

	//如果有错误 就执行panic 中断下面执行的语句
	if err != nil {
		//panic 抛出异常，程序不会再往下执行
		panic(err)
	}
	println("没有执行panic")
}

func test(num1 int, num2 int) int {
	//使用defer和recover 处理异常
	defer func() {
		//捕获错误 赋值给err
		err := recover()
		if err != nil {
			fmt.Println("错误已经被捕获，错误是err", err)
		}
	}()
	return num1 / num2
}

//自定义异常
func test1(num1 int, num2 int) (result int, err error) {
	if num2 == 0 {
		return 0, errors.New("除数不能为0")
	}
	return num1 / num2, nil
}
