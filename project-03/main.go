package main

import (
	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"username" json:"username" xml:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" uri:"password" binding:"required"`
}

func main() {
	//创建路由
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"title": "我的标题"})
	})

	//监听端口，默认8080
	r.Run("0.0.0.0:8000")
}
