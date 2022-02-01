package main

import "fmt"

func main() {
	s := Student{Age: 12}
	p := Person{Age: 30}
	s = Student(p)
	fmt.Println(s)
	fmt.Println(p)
}

type Student struct {
	Age int
}
type Person struct {
	Age int
}
