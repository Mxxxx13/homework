package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"homework7/dao"
	"strconv"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id int
	AccountBalance float64
}

//储存用户注册的信息
func Register(username string, password string) bool{
	stmt, err := dao.DB.Prepare("insert into user(username,password) value (?,?)")
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil {
		fmt.Printf("insert failed:%v", err)
		return false
	}
	return true
}

//检索数据库中的信息并检查密码是否匹配
func Login(username string, password string) bool{
	var u user
	stmt, err := dao.DB.Query("select * from user where username = ?", username)
	if err != nil {
		fmt.Printf("query failed:%v", err)
		return false
	}
	defer stmt.Close()
	for stmt.Next() {
		err = stmt.Scan(&u.Id,&u.Username, &u.Password)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return false
		}
	}
	if u.Password == password {
		return true
	}
	return false
}

//查询账户余额并加上充值金额后更新数据库
func Recharge(username string,rechargeAmount string) bool{
	var u user
	stmt, err := dao.DB.Query("select * from user where username = ?", username)
	if err != nil {
		fmt.Printf("query failed:%v", err)
		return false
	}
	defer stmt.Close()
	for stmt.Next() {
		err = stmt.Scan(&u.Id, &u.Username, &u.AccountBalance)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return false
		}
	}
	RechargeAmount,_:= strconv.ParseFloat(rechargeAmount,64)
	accountBalance := u.AccountBalance + RechargeAmount
	stmt2, err := dao.DB.Prepare("update user set accountBalance = ? where username = ? ")
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt2.Exec(accountBalance,username)
	if err != nil {
		fmt.Printf("insert failed:%v", err)
		return false
	}
	return true
}