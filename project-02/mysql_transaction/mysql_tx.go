package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type user struct {
	id   int64
	name string
	age  int
}

// Go 连接mysql
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库连接失败,err:", err)
	}
	fmt.Println("数据库连接成功")
	transactionDemo()
}

//初始化数据库连接
func initDB() (err error) {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/golang"
	//连接数据库
	db, err = sql.Open("mysql.sql", dsn)
	return
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed,err:%v", err)
		return
	}
	//执行多个事务操作
	sqlStr1 := "update user set age =age -2 where id =1"
	sqlStr2 := "update user1 set age =age +2 where id =1"
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql1出错了，回滚")
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql2出错了，回滚")
		return
	}
	tx.Commit()
	fmt.Println("事务执行成功")
}
