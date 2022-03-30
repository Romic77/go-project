package main

import (
	"blogger/controller"
	"blogger/dao/db"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("blogger")

func main() {
	r := gin.Default()
	driverSql := "root:rootroot@tcp(127.0.0.1:3306)/blogger?parseTime=true&loc=Local"
	db.InitDB(driverSql)

	//加载静态文件
	r.Static("/static/", "./static")
	//加载模板
	r.LoadHTMLGlob("views/*")

	// 首页
	r.GET("/", controller.IndexHandler)

	//目前只完成了首页功能，其他页面参考 https://github.com/pingguoxueyuan/gostudy/blob/9e3f839c61/blogger/controller/handler.go
	r.Run("0.0.0.0:8000")
}
