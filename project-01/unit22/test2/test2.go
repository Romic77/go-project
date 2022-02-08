package main

import (
	"encoding/json"
	"fmt"
)

//json格式Name转义为name
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{Name: "zhangsan", Age: 100}
	//序列化p1
	marshal, err := json.Marshal(p1)
	if err != nil {
		println("marshal error ;%v", err)
		return
	}
	fmt.Printf("%v\n", string(marshal))
	//反序列化
	var p2 person
	json.Unmarshal(marshal, &p2)
	fmt.Printf("%#v\n", p2)
}
