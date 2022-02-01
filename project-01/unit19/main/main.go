package main

func main() {
	var arr = [6]int{3, 6, 9, 1, 2, 4}
	//定义一个切片，名字为slice,是从arr中索引1-3(包含1，不包含3)
	//var slice []int = arr[1:3]

	println("arr下标为1的地址：", &arr[1])
	slice := arr[1:3]

	println("slice下标为0的地址：", &slice[0])

	slice[0] = 123
	println("slice下标为0的地址：", slice[0])
	println("arr下标为1的地址：", arr[1])
	/*for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	println("容量是：", cap(slice))

	for index, value := range slice {
		println("索引是:", index, "；值为：", value)
	}*/
}
