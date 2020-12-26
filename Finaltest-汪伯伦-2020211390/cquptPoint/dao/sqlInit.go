package dao

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func MysqlInit() *sql.DB{
	db, err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8")
	if err != nil {
		fmt.Printf("mysql connect failed: %v\n", err)
		return nil
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Printf("mysql connecr failed: %v\n", err)
		return nil
	}
	DB = db
	return DB
}