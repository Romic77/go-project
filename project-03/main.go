package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"username" json:"username" xml:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" uri:"password" binding:"required"`
}

func main() {
	//创建路由
	r := gin.Default()

	//JSON绑定
	r.GET("/:username/:password", func(c *gin.Context) {
		var uriBind Login
		err := c.ShouldBindUri(&uriBind)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if uriBind.User == "root" && uriBind.Password == "123456" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "ok", "msg": "登录成功！"})
			return
		}
	})
	//监听端口，默认8080
	r.Run("0.0.0.0:8000")
}
