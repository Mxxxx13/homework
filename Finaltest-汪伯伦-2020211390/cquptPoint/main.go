package main

import (
	"homework7/cmd"
	"homework7/dao"
)

func main (){
	dao.MysqlInit()
	cmd.Entrance()

}