package main

import (
	"fmt"
	"go-project/project-01/unit12/dbutils"
)

func main() {
	dbutils.GetConn()
	fmt.Println("hello")
}
