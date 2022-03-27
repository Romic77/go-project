package model

import "time"

// Article
// @description 文章
type Article struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Title        string    `db:"title"`
	Summary      string    `db:"summary"`
	ViewCount    uint32    `db:"view_count"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
}

// ArticleDetail
// @description 文章详情
type ArticleDetail struct {
	Article
	//文章内容
	Content string `db:"content"`
	Category
}

// ArticleRecord
// @description 文章上下页
type ArticleRecord struct {
	Article
	Category
}
