package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	name := flag.String("name", "张三", "请输入名称")
	age := flag.Int("age", 9000, "请输入真实年龄")
	married := flag.Bool("married", false, "结婚了吗")
	duration := flag.Duration("time", time.Second, "请输入时间")

	/*var name string
	flag.StringVar(&name, "name", "张三", "请输入名称")
		flag.IntVar(&age, "age", 9000, "请输入真实年龄")
		flag.BoolVar(&married, "married", false, "结婚了吗")
		flag.DurationVar(&time, "time", time.Second, "请输入时间")
	*/
	//使用flag
	flag.Parse()
	fmt.Println(name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*duration)
	//
	fmt.Println(flag.Args())  //返回命令行参数后的其他参数
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数的个数  以-个数
}
