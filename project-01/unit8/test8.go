package main

import "fmt"

func main() {
	var str string = "hello 你好"

lable1:
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
		for i2 := 0; i2 < 10; i2++ {
			if i2 == 1 {
				break lable1
			}
		}
	}

	for i, value := range str {
		fmt.Printf("索引是：%d,值是：%c \n", i, value)
	}

}
