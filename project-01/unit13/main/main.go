package main

import "go-project/project-01/unit13/test"

func init() {
	println("main init")
}

func main() {
	println(test.GlobleName)
	println("hello world")

	test.Add()

	//先执行init方法，在执行全局变量，最后执行函数
}
