package main

import "fmt"

type student struct {
	id   int
	name string
}

//学生管理结构体
type stuManager struct {
	allStudent map[int]student
}

//添加学生
func (s stuManager) addStudent() {
	var s1 student
	fmt.Println("请输入学生的学号：")
	fmt.Scanln(&s1.id)
	fmt.Println("请输入学生的姓名：")
	fmt.Scanln(&s1.name)
	s.allStudent[s1.id] = s1
}

//删除学生
func (s stuManager) deleteStudent() {
	var s1 student
	fmt.Println("请输入需要删除的学号")
	fmt.Scanln(&s1.id)
	delete(s.allStudent, s1.id)
}

//修改学生
func (s stuManager) updateStudent() {
	var s1 student
	fmt.Println("请输入学生的学号：")
	fmt.Scanln(&s1.id)
	fmt.Println("请输入学生的姓名：")
	fmt.Scanln(&s1.name)
	stuObj, ok := s.allStudent[s1.id]
	if !ok {
		println("查无此人")
		return
	}
	println("你要修改的学生信息如下：学号:", s1.id, "姓名为：", stuObj.name)
	s.allStudent[s1.id] = stuObj

}

//查询学生
func (s stuManager) queryStudents() {
	for key, value := range smr.allStudent {
		println("学生id是：", key, "学生的名字是:", value.name)
	}
}
