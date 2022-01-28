package main

import "fmt"

func main() {
	//test()
	//test1()
	//test2()
	//test3()
	//test4()
	test6()
}

func test() {
	//5个学生的总成绩和平均分输
	score1 := 90
	score2 := 91
	score3 := 92
	score4 := 93
	score5 := 94

	sum := score1 + score2 + score3 + score4 + score5
	avg := sum / 5

	println("总成绩是:", sum, "平均分是:", avg)
}

func test1() {
	var scores = [5]int{90, 91, 92, 93, 94}

	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	avg := sum / len(scores)
	println("总成绩是:", sum, "平均分是:", avg)
}

func test2() {
	var scores = [5]int{}
	sum := 0
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请录入第%d个学生的成绩", i+1)
		//控制台输入值需要加上&
		fmt.Scanln(&scores[i])
		sum += scores[i]
	}
	avg := sum / len(scores)
	println("总成绩是:", sum, "平均分是:", avg)

	for i, score := range scores {
		println("第:", i+1, "个人的分数是", score)
	}
}

func test3() {
	var scores = [...]int{2: 66, 0: 33, 1: 99}
	println(scores[1])

	var scores1 = [...]int{}
	for i := 0; i < len(scores1); i++ {
		scores1[i] = i + 1
	}
}

func test4() {
	var scores = [...]int{1, 2, 3}
	test5(&scores)
	fmt.Println(scores)
}

//把scores当做指针传递过来， 然后根据指针修改真实值
func test5(scores *[3]int) {
	//修改指针为0的值
	(*scores)[0] = 6
}

//二维数组 是数组里面包含数组
func test6() {
	var arr = [2][3]int16{}
	fmt.Println(arr)

	fmt.Printf("arr[0]的地址是：%p\n", &arr[0])
	fmt.Printf("arr[0][0]的地址是：%p\n", &arr[0][0])

	//赋值
	arr[0][1] = 40
	arr[0][0] = 91

	//fmt.Println(arr)
	/*for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			println(arr[i][j])
		}
	}*/

	//定义2行3列的二维数组
	//1 ，2 ，0 ,0
	//3 , 4, 5 ,0
	//6 , 7, 8 ,0
	var scores = [3][4]int{{1, 2}, {3, 4, 5}, {6, 7, 8}}
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			println(scores[i][j])
		}
		fmt.Println()
	}
	//fmt.Println(scores)

	for key, value := range scores {
		for k, v := range value {
			fmt.Printf("arr[%v][%v]=%v\t", key, k, v)
		}
		fmt.Println()
	}
}
