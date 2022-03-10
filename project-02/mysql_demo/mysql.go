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
	//queryById(1)
	queryParams()
	//初始化结构体的方式 new -map 切片 通道,new函数返回的是user的内存地址(指针)
	/*user := new(user)
	user.age = 12
	user.name = "new zhangsan"*/

	//结构体是基本数据类型， 使用&或者new初始化数据 就是指针类型。 否知就是值类型

	u := user{
		name: "zhangsan2",
		age:  12,
	}
	u.id = 4
	u.name = "张三3"
	u.age = 50
	//insertUser(u)
	updateUser(u)
	deleteById(u)
}

//初始化数据库连接
func initDB() (err error) {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/golang"
	//连接数据库
	db, err = sql.Open("mysql", dsn)
	return
}

//数据库根据id查询
func queryById(id int) {
	var user user

	// 1.写查询单挑记录的sql语句
	sqlStr := "select * from user where id =?;"
	// 2.执行
	rowObj := db.QueryRow(sqlStr, id)

	//scan会释放数据库的链接
	rowObj.Scan(&user.id, &user.name, &user.age)
	fmt.Println(user)
}

func queryParams() {
	var user user

	sqlStr := "select * from user where id > ?;"
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	rowObj, err := stmt.Query(sqlStr, 0)
	if err != nil {
		fmt.Println("exec %s query failed,error: %v", sqlStr, err)
	}
	defer rowObj.Close()
	for rowObj.Next() {
		rowObj.Scan(&user.id, &user.name, &user.age)
		fmt.Println(user)
	}
}

//insert 就传递值
func insertUser(u user) {
	sqlStr := "insert into user(name,age) values(?,?);"
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()

	ret, _ := stmt.Exec(sqlStr, u.name, u.age)
	//插入数据，拿到数据的id
	id, _ := ret.LastInsertId()
	fmt.Printf("got LastInsertId: %v", id)
}

func updateUser(u user) {
	sqlStr := "update user set name=? , age = ? where id = ?;"
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	ret, _ := stmt.Exec(sqlStr, u.name, u.age, u.id)
	//返回受影响的行数
	rowsAffected, _ := ret.RowsAffected()
	fmt.Printf("update rowsAffected: %v\n", rowsAffected)
}

func deleteById(u user) {
	sqlStr := "delete from user where id = ?;"
	stmt, _ := db.Prepare(sqlStr)
	defer stmt.Close()
	ret, _ := stmt.Exec(sqlStr, u.id)
	//返回受影响的行数
	rowsAffected, _ := ret.RowsAffected()
	fmt.Printf("delete rowsAffected: %v\n", rowsAffected)
}
