package service

import (
	"github.com/gin-gonic/gin"
	"homework7/models"
)

//获取用户注册的信息
func Register(ctx *gin.Context) bool {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	res := models.Register(username, password)
	return res
}

//获取用户登录的信息
func Login (ctx *gin.Context) bool {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	res := models.Login(username, password)
	return res
}


func Recharge (ctx *gin.Context) bool {
	rechargeAmount := ctx.PostForm("rechargeAmount")
	username := ctx.PostForm("username")
	res := models.Recharge(username,rechargeAmount)
	return res
}