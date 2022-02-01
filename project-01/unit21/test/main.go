package main

import "fmt"

type Student struct {
	Age int
}

type Stu struct {
	Age int
}

func main() {
	var s1 = Student{1}
	var s2 = Stu{20}
	s1 = Student(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}
