package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Login struct {
	User     string `form:"username" json:"username" xml:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" uri:"password" binding:"required"`
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		// 设置变量到Context的key中
		c.Set("request", "中间件")
		//执行中间件
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	//创建路由
	r := gin.Default()

	//注册中间件
	r.Use(middleware())
	{
		r.GET("/middleware", func(c *gin.Context) {
			get, _ := c.Get("request")
			fmt.Println("request:", get)
			//页面接收
			c.JSON(200, gin.H{"request": get})
		})

		//跟路由后面定义的是局部中间件
		r.GET("/middleware2", middleware(), func(c *gin.Context) {
			get, _ := c.Get("request")
			fmt.Println("request:", get)
			//页面接收
			c.JSON(200, gin.H{"request": get})
		})
	}
	//监听端口，默认8080
	r.Run("0.0.0.0:8000")
}
