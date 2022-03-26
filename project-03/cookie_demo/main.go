package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthHandler(c *gin.Context) {
	cookie, err := c.Cookie("abc")
	if cookie == "123" {
		c.Next()
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": err})
	//中断
	c.Abort()
	return
}

func main() {
	//创建路由
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		c.String(200, "login success")
	})

	//访问主页需要cookie的值
	r.GET("/home", AuthHandler, func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})

	r.Run("0.0.0.0:8000")
}
