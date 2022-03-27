package db

import (
	"blogger/model"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("blogger")

// InsertCategory
// @description 添加分类
// @param category *model.Category
// @return categoryId int64
// @return err error
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := "insert into category(category_name,category_no) values(?,?);"
	result, err := db.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

// GetCategoryById
// @description 根据分类id获取分类对象
// @param id int64
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlStr := "select id,category_name,category_no,create_time,update_time from category where id = ?"
	err = db.Get(category, sqlStr, id)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

// GetCategoryListByIds
// @description 根据分类id切片查询分类
// @return categoryList []*model.Category
// @return err error
func GetCategoryListByIds(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlStr := "select * from category where category.id in (?);"
	sqlStr, args, err := sqlx.In(sqlStr, categoryIds)
	err = db.Select(&categoryList, sqlStr, args)
	if err != nil {
		return nil, err
	}
	return
}

// GetCategoryList
// @description 获取所有分类
// @return categoryList []*model.Category
// @return err error
func GetCategoryList() (categoryList []*model.Category, err error) {
	sqlStr := "select * from category order by category_no asc;"
	//Select方法帮你做了 初始化
	err = db.Select(&categoryList, sqlStr)
	if err != nil {
		return
	}
	return
}
