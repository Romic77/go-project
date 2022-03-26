package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func InitDB() (err error) {
	addr := "root:rootroot@tcp(127.0.0.1:3306)/golang"
	db, err = sqlx.Open("mysql", addr)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}

// queryBookList
// @description 查询所有书籍
// @return bookList *[]Book
// @return err error
func queryBookList() (bookList []*Book, err error) {
	sqlStr := "select * from book;"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	return
}

// InsertBook
// @description 插入书籍
// @param book *Book
// @return err error
func InsertBook(book *Book) (err error) {
	sqlStr := "insert into book(title,price) values(?,?);"
	_, err = db.Exec(sqlStr, book.Title, book.Price)
	if err != nil {
		return err
	}
	return
}

// updateBook
// @description 更新书籍信息
// @param book *Book
// @return err error
func updateBook(book *Book) (err error) {
	sqlStr := "update book set title =? and price = ? where id =?;"
	_, err = db.Exec(sqlStr, book.Title, book.Price, book.Id)
	if err != nil {
		return err
	}
	return
}

// DeleteBookById
// @description 删除书记根据id
// @param book *Book
// @return err error
func DeleteBookById(id int) (err error) {
	sqlStr := "delete from book where id =?;"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return
}
