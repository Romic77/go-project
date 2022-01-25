package main

func main() {
	var age int = 32

	if age > 20 {
		println("老人")
	} else {
		println("小孩")
	}

	switch age {
	case 32:
		println("老人1")
		//switch 穿透
		fallthrough
	case 31:
		println("老人2")

	default:
		println("小孩")
	}
}
