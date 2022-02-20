package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:name`
	Age  int    `json:age`
}

func main() {
	str := `{"name":"xx","age":100}`
	var p person
	//通过方法的反射实现的
	json.Unmarshal([]byte(str), &p)

	fmt.Println(p.Name, p.Age)

	var a float32 = 3.14
	reflectType(a)

	var b int = 100
	reflectType(b)

	reflectValue(a)
	reflectValue(b)

	//reflectSetValue1(&b)
	reflectSetValue2(&b)
	fmt.Println("修改后的值是", b)

	stu1 := student{
		Name:  "张三",
		Score: 50,
	}
	//获取t的类型
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name %s index:%d type:%v json tag :%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name %s index:%d type:%v json tag :%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name :%v type kind :%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	//获取x的值
	v := reflect.ValueOf(x)
	//获取x的值的类型
	k := v.Kind()
	switch k {
	case reflect.Int:
		fmt.Printf("type is int , value is %v\n", int(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is int , value is %f\n", float32(v.Float()))
	}
}

//通过反射修改值,只能传递指针
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int {
		v.SetInt(200)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(200)
	}
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
