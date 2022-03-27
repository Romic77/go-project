package db

import (
	"testing"
)

//init 是单元测试默认就会前置执行的函数
func init() {
	//parseTime=true 将MYSQL中的时间类型 自动解析为go结构体中的时间类型
	driverStr := "root:rootroot@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	InitDB(driverStr)
}

// TestGetCategoryById
// @description 测试GetCategoryById
// @param t *testing.T
func TestGetCategoryById(t *testing.T) {
	category, _ := GetCategoryById(1)
	if category == nil {
		t.Fatal("失败")
	}
	t.Logf("category: %#v", category)
}

func TestGetCategoryList(t *testing.T) {
	categoryList, _ := GetCategoryList()
	if len(categoryList) == 0 {
		t.Fatal("categoryList 长度为0")
	}

	for _, category := range categoryList {
		t.Logf("categoryList: %#v\n", category)
	}
}
