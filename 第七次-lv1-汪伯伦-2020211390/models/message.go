package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"homework7/dao"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

//将留言储存到数据库内
func SaveMessage(username string, message string) bool {
	stmt, err := dao.DB.Prepare("insert into messageBoard(username,message)values(?,?)")
	if err != nil {
		fmt.Printf("prepare failed: %v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, message)
	if err != nil {
		fmt.Printf("insert failed: %v", err)
		return false
	}
	return true
}

//检索数据库中的留言
func ViewMessage(username string) []Message {
	stmt, err := dao.DB.Query("select * from messageBoard where username = ?", username)
	if err != nil {
		fmt.Printf("query failed: %v", err)
		return nil
	}
	defer stmt.Close()
	var msgs []Message
	for stmt.Next() {
		var m Message
		_ = stmt.Scan(&m.Username, &m.Message)
		msgs = append(msgs, m)
	}
	return msgs
}