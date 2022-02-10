package main

import "fmt"

func main() {
	//inferface{} 是空接口
	// 声明一个stududentMap,key为string value是空接口（类似泛型）
	var studentMap map[string]interface{}

	studentMap = make(map[string]interface{}, 16)
	studentMap["name"] = "张三"
	studentMap["age"] = 123

	fmt.Print(studentMap)

	show(false)
	show("123")
	show(nil)
	show(studentMap)

	//接口.(类型) -> 意思就是把类型强制转为类型
	assign(1)

	assign2(1)
	assign2("true")
}

//参数类型为空接口。类似泛型
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

//空接口可以传递任意类型，那么怎么判断是什么类型呢。java是通过instanceof
func assign(a interface{}) {
	str, ok := a.(string)
	if !ok {
		fmt.Printf("不是string类型")
	}
	fmt.Printf(str)
}

func assign2(a interface{}) {
	switch t := a.(type) {
	case string:
		fmt.Println("是一个string", t)
	case int:
		fmt.Println("是一个int", t)
	case bool:
		fmt.Println("是一个bool", t)
	default:
		fmt.Println("不知道")
	}
}
