package main

func main() {
	//new(类型) 返回对应类型的指针
	num := new(int)
	println(num)
	println(*num)
}
