package main

import "fmt"

func main() {
	var age int
	println("请录入学生的年龄：")
	fmt.Scanln(&age)

	var name string
	println("请录入学生的姓名：")
	fmt.Scanln(&name)

	var score float32
	println("请录入学生的成绩：")
	fmt.Scanln(&score)

	var isVIP bool
	println("请录入学生是否为VIP：")
	fmt.Scanln(&isVIP)

	println("学生的年龄是：%v,姓名是%v，成绩是%v，是否为vip%v", age, name, score, isVIP)
}
