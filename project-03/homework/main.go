package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// main
// @Description: 定义一个中间件打印每个url执行时间
func main() {
	r := gin.Default()

	r.Use(costTime)
	{
		group := r.Group("/shopping")
		{
			group.GET("/index", shopIndexHandler)
			group.GET("/home", shopHomeHandler)
		}
	}

	r.Run("0.0.0.0:8000")
}

func costTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	//统计时间
	since := time.Since(start)
	fmt.Println("程序用时:", since)
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(1 * time.Second)
}
