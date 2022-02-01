package main

import "fmt"

func main() {
	//make函数的三个参数：切片类型、切片长度、切片容量
	//make底层创建了一个数组，但是这个数组无法进行操作，需要通过slice切片进行见解访问。
	/*slice := make([]int, 4, 20)

	slice[0] = 3
	slice[1] = 3
	fmt.Println(slice)

	slice2 := []int{1, 4, 5}
	fmt.Println(slice2)*/

	//test1()
	test2()
}

func test1() {
	slice := []int{}
	fmt.Println(slice)

	//定义数组
	intArr := [6]int{1, 4, 7, 2, 5, 8}
	//对intArr进行切片,从索引为1的位置开始切，一直到索引为4但是不包含4
	//[4 7 2]
	slice1 := intArr[1:4]
	fmt.Println(slice1)

	//对slice1再次进行切片,从索引为1的位置进行切片，一直到索引为2
	//[7]
	slice2 := slice1[1:2]
	fmt.Println("slice2的切片值为：", slice2)
	slice2[0] = 99
	//修改slice[0]的值为99之后，原来的intArr的值也进行了改变，把intArry[2]的值改为了99
	//因为切片是三个关健，切片地址，切片长度，切片容量
	fmt.Println(intArr)
}

//切片的动态增长
func test2() {
	intArr := [6]int{1, 4, 5, 6, 7, 8}
	slice := intArr[1:4]
	fmt.Println("slice的长度为：", len(slice))
	//在slice切片中添加2个数组，88，50
	//slice2的值为 [4 5 6 88 50]
	slice2 := append(slice, 88, 50)
	fmt.Println(slice2)

	//append的底层原理
	//1 底层追加元素的时候对数组进行扩容，老数组扩容为新数组
	//2 创建1个新数组，将老数组中的4 5 6 复制到新数组中去 在新数组中追加88 50
	//3 slice2底层数组的指向的是新数组

	//slice 和slice2 合并 赋值给slice3
	slice3 := append(slice, slice2...)
	fmt.Println(slice3)
}
