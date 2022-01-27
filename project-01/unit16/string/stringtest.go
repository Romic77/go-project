package main

import (
	"fmt"
	strconv "strconv"
	"strings"
)

func main() {
	str := "hello 你好"
	println(len(str))

	//for-range 遍历
	for i, value := range str {
		fmt.Printf("%d: %c\n", i, value)
	}

	//切片遍历
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%d: %c\n", i, r[i])
	}

	//字符串转整数
	num, _ := strconv.Atoi("123")
	println(num)

	var str1 string = strconv.Itoa(321)
	println(str1)

	//查询字符串中l出现的次数
	println(strings.Count(str, "l"))

	//EqualFold 字符串不区分大小写
	println(strings.EqualFold("go", "Go"))

	println("Go" == "go")

	//查找 o是否在字符串str中，存在返回索引下标，不存在返回-1
	println(strings.Index(str, "o"))

	//替换str字符串中l，把l替换成O,替换1次
	println(strings.Replace(str, "l", "O", 1))

	arr := strings.Split(str, "l")
	fmt.Println(arr)
	println(strings.ToUpper(str))
	println(strings.ToLower(str))

	println(strings.TrimSpace(str))

	//去除左边的hello
	println(strings.TrimLeft(str, "hello"))
	//去除右边的好
	println(strings.TrimRight(str, "好"))

	//以h开头
	println(strings.HasPrefix(str, "h"))

	//以h结尾
	println(strings.HasSuffix(str, "h"))

}
