package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Print("欢迎光临学生管理系统")
	fmt.Println(`
	1.查看所有学生
	2.新增学生
	3.修改学生
	4.删除学生
	5.退出`)
}

//全局的学生变量
var smr = stuManager{allStudent: make(map[int]student)}

// 学生管理系统
func main() {
	for {
		showMenu()
		fmt.Print("请输入序号：")
		//等待用户输入
		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			smr.queryStudents()
		case 2:
			smr.addStudent()
		case 3:
			smr.updateStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("GG")
		}
	}
}
