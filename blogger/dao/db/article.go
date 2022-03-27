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

// GerArticleList
// @description 获取文章列表
// @param pageNum int
// @param pageSize int
func GerArticleList(pageNum int, pageSize int) {

}
