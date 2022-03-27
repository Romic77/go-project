package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

// InitDB
// @description 初始化MySQL数据库
// @param driverStr string
// @return err error
func InitDB(driverStr string) (err error) {
	db, err = sqlx.Open("mysql", driverStr)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(10)
	return
}
