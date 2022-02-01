package main

import "fmt"

func main() {
	//定义map变量,只声明 是不分配内存空间
	var a map[int]string
	//必须通过make函数进行初始化，才会分配内存空间
	a = make(map[int]string, 10)
	//将键值对存入map
	a[1] = "张三"
	a[2] = "李四"
	a[3] = "王五"

	//fmt.Println(a)

	//test1()
	//test2()
	test3()
}

func test1() {
	//创建map的方式2
	b := make(map[int]string, 10)
	b[1] = "张三"
	b[2] = "李四"
	b[3] = "王五"

	fmt.Println(b)

	//创建map的方式3
	c := map[int]string{1: "张三", 2: "李四", 3: "王五", 4: "赵六"}
	fmt.Println(c)
}

func test2() {
	b := make(map[int]string, 10)
	b[1] = "张三"
	b[2] = "李四"
	b[3] = "王五"
	//覆盖key=3的值为赵六
	b[3] = "赵六"
	//删除b中key=3的
	delete(b, 3)
	fmt.Println(b)
	//获取b[2]的值
	value, flag := b[2]
	fmt.Println(value, flag)

	//清空b的所有值，重新构造一个b
	b = make(map[int]string, 10)
	fmt.Println(b)
}

//map遍历
func test3() {
	b := make(map[int]string, 10)
	b[1] = "张三"
	b[2] = "李四"
	b[3] = "王五"
	/*for key, value := range b {
		println(key, value)
	}*/

	c := make(map[int]string, 10)
	c[1] = "张三"

	//a key是string value是map
	a := make(map[string]map[int]string, 10)
	a["班级1"] = b
	a["班级2"] = c

	for key, value := range a {
		println(key)
		//再次遍历value 为map
		for k1, v1 := range value {
			println(k1, v1)
		}

	}
}
