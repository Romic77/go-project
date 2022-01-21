package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

var age = 0

func main() {
	//var age = 123
	var name = "chenqi"
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(unsafe.Sizeof(age))
	var name1 int = '中'
	fmt.Println(name1)
	fmt.Printf("%c\n", name1)

	//双引号转义
	fmt.Println("\"golang\"")
	//反斜杠转义
	fmt.Println("\\golang\\")
	//换行
	fmt.Println("hello\n")
	//\r 贯标回到本行开头，后续输入的字符替换原来的字符
	fmt.Println("hello\rworld")

	var flag bool = true
	fmt.Println(flag)

	var str string = "sb "
	str = "da sb"
	fmt.Println(str)

	var str1 string = `//\r 贯标回到本行开头，后续输入的字符替换原来的字符
	fmt.Println("hello\rworld")

	var flag bool = true
	fmt.Println(flag)

	var str string = "sb "
	str = "da sb"
	fmt.Println(str)`
	println(str1)

	var floatA float64 = 21.611
	var a int = int(floatA)
	println(a)

	var n1 int = 19
	/*	var n2 float32 = 4.78
		var n3 bool = false
		var n4 byte = 'a'*/
	var s1 string = fmt.Sprintf("%d", n1)
	println(s1)

	strconv.FormatInt(int64(n1), 10)
}
