package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	//初始化数据库
	err := InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()
	//加载页面
	r.LoadHTMLGlob("./book/templates/*")
	//查询所有图书
	r.GET("/book/list", bookListHandler)

	//页面跳转
	r.GET("/book/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new_book.html", gin.H{"code": 200})
	})

	//添加图书
	r.POST("/book/new", bookAddHandler)

	//删除图书
	r.GET("/book/delete", bookDeleteHandler)

	r.Run("0.0.0.0:8000")
}

func bookListHandler(c *gin.Context) {
	list, err := queryBookList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err})
		return
	}
	c.HTML(http.StatusOK, "book_list.html", gin.H{"code": 200, "data": list})
}

func bookAddHandler(c *gin.Context) {
	var book Book
	c.Bind(&book)
	err := InsertBook(&book)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func bookDeleteHandler(c *gin.Context) {
	param := c.Query("Id")
	id, _ := strconv.Atoi(param)
	fmt.Println(id)
	DeleteBookById(id)
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}
