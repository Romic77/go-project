package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	//1. 从数据库中，获取文章分类列表
	categoryList, err = db.GetCategoryList()
	if err != nil {
		fmt.Printf("1 get article list failed, err:%v\n", err)
		return
	}

	return
}
