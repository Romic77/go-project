package main

import (
	"fmt"
	"os"
)

//函数版学生管理系统（查看学生、新增学生、删除学生）
type Student struct {
	id   int64
	name string
}

func main() {
	allStudent = make(map[int64]Student, 50)
	for {
		//1.打印菜单
		//2.等待用户选择
		//3.执行对应的函数
		fmt.Print("欢迎光临学生管理系统")
		fmt.Println(`
	1.查看所有学生
	2.新增学生
	3.删除学生
	4.退出`)
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("您选择了%d这个选项\n", choice)
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("GG")
		}
	}
}

var (
	allStudent map[int64]Student //变量声明
)

func deleteStudent() {
	fmt.Println("请输入需要删除的学号")
	var (
		id int64
	)
	fmt.Scanln(&id)
	delete(allStudent, id)
	fmt.Println("删除成功，学号为:", id)
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Println("请输入学生的学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生的姓名：")
	fmt.Scanln(&name)
	s1 := Student{id, name}
	allStudent[id] = s1
}

func showAllStudent() {
	for key, value := range allStudent {
		fmt.Println("学号为：", key, ";  姓名是：", value.name)
	}

}
