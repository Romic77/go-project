package controller

import (
	"blogger/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//
// IndexHandler
// @Description 主页处理器
//
func IndexHandler(c *gin.Context) {
	//从service中取数据
	//加载文章数据
	articleList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{"article_list": articleList, "category_list": categoryList})
}

//
// CategoryList
// @Description 点击分类云进行分类
//
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	articleList, err := service.GetArticleListById(categoryId, 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{"article_list": articleList, "category_list": categoryList})
}
