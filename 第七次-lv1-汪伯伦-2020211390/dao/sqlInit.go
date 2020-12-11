package dao

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func MysqlInit() *sql.DB{
	db, err := sql.Open("mysql","root:20030513wbl@tcp(127.0.0.1:3306)/sql_test?charset=utf8")
	if err != nil {
		fmt.Println("mysql connect failed: %v", err)
		return nil
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Println("mysql connecr failed: %v", err)
		return nil
	}
	DB = db
	return DB
}