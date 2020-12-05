//操作数组库：创建“uesr”表
//CREATE TABLE `user` (
//`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
//`name` VARCHAR(20) DEFAULT '',
//`email` varchar(255) DEFAULT '0',
//password varchar(255) DEFAULT '123456',
//PRIMARY KEY(`id`)
//)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

type user struct {
	id   int
	name string
	password string
	email string
}

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:******@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Println("init db failed,err:%v", err)
		return
	}else{
		fmt.Println("init db success")
	}

	//增
	Insert :="insert into user(name,email,password) value(?,?,?)"
	_, err = db.Exec(Insert, "zhangsan", "xxxxxxxx@qq.com","******")
	if err != nil {
		fmt.Println("insert failed, err:%v", err)
		return
	}else{
		fmt.Println("insert success")
	}

	// 删
	Delete :="delete from user where id = ?"
	_,err = db.Exec(Delete,1)
	if err != nil {
		fmt.Println("delete failed, err:%v\n", err)
	} else {
		fmt.Println("delete success")
	}

	//改
	Update := "update user set password = ? where id = ?"
	_,err = db.Exec(Update,"123456",1)
	if err != nil {
		fmt.Println("update failed, err:%v",err)
	} else {
		fmt.Println("update success")
	}

	//查
	Select := "select id, name from user where id = ?"
	var u user
	err = db.QueryRow(Select, 1).Scan(&u.id, &u.name)
	if err != nil {
		fmt.Println("select failed, err:%v",err)
	} else {
		fmt.Println("select success")
	}
}
