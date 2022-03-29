package service

import (
	"blogger/dao/db"
	"blogger/model"
)

//
// GetArticleRecordList
// @Description: 获取文章和对应的分类
// @return articleRecordList
// @return err
//
func GetArticleRecordList(pageNum int, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1. 获取文章列表
	articleList, err := db.GetArticleList(pageNum, pageSize)

	//2.获取文章对应的分类(多个)
	ids := GetCategoryIds(articleList)
	categoryList, err := db.GetCategoryListByIds(ids)

	//返回页面 做聚合
	for _, article := range articleList {
		articleObj := &model.ArticleRecord{Article: *article}

		categoryId := articleObj.CategoryId

		for _, category := range categoryList {
			if categoryId == category.Id {
				articleObj.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleObj)
	}
	return
}

//
// GetCategoryIds
// @Description: 根据多个文章id查询多个分类id的集合
// @param articleInfoList
// @return ids
//
func GetCategoryIds(articleList []*model.Article) (ids []int64) {
LABEL:
	//遍历文章
	for _, article := range articleList {
		//从当前文章取出分类id
		categoryId := article.CategoryId
		//去重 分类id如果一样只需要加载一次
		for _, id := range ids {
			if id != categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

//
// GetArticleListById
// @Description 根据分类id 获取该类文章对应的分类信息
// @param categoryId int64
// @param pageNum int
// @param pageSize int
// @return articleRecordList *[]model.ArticleRecord
// @return err error
//
func GetArticleListById(categoryId int64, pageNum int, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1. 获取文章列表
	articleList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)

	//2.获取文章对应的分类(多个)
	ids := GetCategoryIds(articleList)
	categoryList, err := db.GetCategoryListByIds(ids)

	//返回页面 做聚合
	for _, article := range articleList {
		articleObj := &model.ArticleRecord{Article: *article}

		categoryId := articleObj.CategoryId

		for _, category := range categoryList {
			if categoryId == category.Id {
				articleObj.Category = *category
				break
			}
		}
	}
	return
}
