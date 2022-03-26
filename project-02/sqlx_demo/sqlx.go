package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

type User struct {
	Id   int64
	Name string
	Age  int
}

// Go 连接mysql
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库连接失败,err:", err)
	}
	fmt.Println("数据库连接成功")
	user := User{Id: 1}
	user = queryById(user)
	fmt.Printf("%#v\n", user)
	userList := queryAll()
	fmt.Printf("%#v\n", userList)
	//结构体如果没有大写，那么make的时候就找不到类型
	//aa := make([]User, 10)
}

//初始化数据库连接
func initDB() (err error) {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/golang"
	//连接数据库
	db, err = sqlx.Connect("mysql.sql", dsn)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

func queryById(user User) User {
	sqlStr := "select id,name,age from User where id = ?;"
	db.Get(&user, sqlStr, user.Id)
	fmt.Printf("%#v\n", user)
	return user
}

func queryAll() []User {
	var userList []User
	sqlStr := "select id,name,age from User;"
	db.Select(&userList, sqlStr)
	fmt.Printf("%#v\n", userList)
	return userList
}
