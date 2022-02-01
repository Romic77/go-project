package main

import "fmt"

type Teacher struct {
	//外界可以访问Name属性
	Name   string
	age    int
	school string
}

func main() {
	var t1 Teacher
	//没有赋值的时候，有默认值 空 0 空{ 0 }
	fmt.Println(t1)
	t1.Name = "mic"
	t1.school = "gupao"
	t1.age = 31
	fmt.Println(t1)

	t2 := Teacher{"tom", 35, "中南财大"}
	fmt.Println(t2)

	//指针方式创建对象
	t3 := new(Teacher)
	//本来是需要使用(*t3).Name进行赋值.Go对这种方式进行了简化更加符合程序猿编程习惯，直接使用对象.属性进行赋值，编译器还是会转化为(*t3).Name
	(*t3).Name = "原始赋值方式"
	t3.Name = "james"
	t3.age = 50
	t3.school = "小学毕业"
	//获取t3的真实值
	fmt.Println(*t3)

	//指针方式创建对象，使用&符号修改地址
	t4 := &Teacher{"陈奇", 30, "高中毕业"}
	fmt.Println(*t4)
}
