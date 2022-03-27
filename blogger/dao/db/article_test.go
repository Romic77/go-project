package db

import (
	"blogger/model"
	"fmt"
	"testing"
	"time"
)

//init 是单元测试默认就会前置执行的函数
func init() {
	//parseTime=true 将MYSQL中的时间类型 自动解析为go结构体中的时间类型
	driverStr := "root:rootroot@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	InitDB(driverStr)
}

func TestInsertArticle(t *testing.T) {
	a := &model.ArticleDetail{
		Article: model.Article{
			CategoryId:   1,
			Title:        "测试",
			Summary:      "测试摘要",
			ViewCount:    0,
			CommentCount: 0,
			Username:     "username",
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
		},
		Content: "",
		Category: model.Category{
			Id:           1,
			CategoryName: "",
			CategoryNo:   0,
			CreateTime:   time.Time{},
			UpdateTime:   time.Time{},
		},
	}

	articleId, _ := InsertArticle(a)
	if articleId == 0 {
		t.Fatal("插入失败")
	}
	fmt.Println("插入返回主键", articleId)
}
