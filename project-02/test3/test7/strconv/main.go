package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "100"
	parseInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	fmt.Println(strconv.Atoi("12"))

	fmt.Println(strconv.Itoa(12))

	fmt.Println(strconv.ParseBool("false"))

	fmt.Println(strconv.ParseFloat("3.1415", 64))
}
