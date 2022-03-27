package db

import (
	"blogger/model"
)

// InsertArticle
// @description 插入文章
func InsertArticle(articleDetail *model.ArticleDetail) (id int64, err error) {
	sqlStr := `insert into article(category_id,content,title,view_count,comment_count,username,summary,create_time,update_time) 
               values(?,?,?,?,?,?,?,?,?)`
	result, err := db.Exec(sqlStr, articleDetail.Article.CategoryId, articleDetail.Content, articleDetail.Article.Title, articleDetail.Article.ViewCount, articleDetail.Article.CommentCount,
		articleDetail.Article.Username, articleDetail.Article.Summary, articleDetail.Article.CreateTime, articleDetail.Article.UpdateTime)
	id, _ = result.LastInsertId()
	return
}

// GetArticleList
// @description 根据分页获取文章列表
// @param pageNum int
// @param pageSize int
// @return articleList []*model.Article articleList指针的值
// @return err error
func GetArticleList(pageNum int, pageSize int) (articleList []*model.Article, err error) {
	if pageNum < 0 || pageSize <= 0 {
		log.Info("pageNum:{},pageSize:{}", pageNum, pageSize)
		return
	}

	sqlStr := `SELECT id,summary,title,view_count,comment_count,username,category_id,create_time FROM article 
				where status = 1 order by create_time desc
				limit ?,?`
	err = db.Select(&articleList, sqlStr, pageNum-1, pageSize)
	if err != nil {
		return nil, err
	}
	return
}

// getArticleById
// @description 根据文章id获取文章对象
// @param id int64
// @return article *model.Article
// @return err error
func GetArticleById(id int64) (article *model.Article, err error) {
	sqlStr := "select id,summary,title,view_count,comment_count,username,category_id,create_time FROM article  where id = ? and status = 1"
	db.Get(&article, sqlStr, id)
	return
}

// GetArticleListByCategoryId
// @description 根据分类id查询文章列表
// @param categoryId int64
// @param pageNum int
// @param pageSize int
// @return articleList *[]model.Article
// @return err error
func GetArticleListByCategoryId(categoryId int64, pageNum int, pageSize int) (articleList *[]model.Article, err error) {
	if pageNum < 0 || pageSize <= 0 {
		log.Info("pageNum:{},pageSize:{}", pageNum, pageSize)
		return
	}
	sqlStr := `id,summary,title,view_count,comment_count,username,category_id,create_time FROM article where status =1 and category_id =?
		order by create_time desc limit ?,?`
	db.Select(&articleList, sqlStr, categoryId, pageNum, pageSize)
	return
}
