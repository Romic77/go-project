package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	// 初始化 es链接
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9201"))
	if err != nil {
		panic(err)
	}

	fmt.Println("connect es success")

	s1 := Student{
		Name:    "romic",
		Age:     30,
		Married: true,
	}
	put, err := client.Index().Index("student").BodyJson(s1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Index user %s to index %s,type %s\n", put.Id, put.Index, put.Type)
}
