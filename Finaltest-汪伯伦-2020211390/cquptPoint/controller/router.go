package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homework7/service"
	"net/http"
)

//返回注册后的信息
func Register(ctx *gin.Context){
	res := service.Register(ctx)
	if res {
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "register success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "register failed",
		})
	}
}

//返回登录后的信息并设置cookie
func Login(ctx *gin.Context){
	res := service.Login(ctx)
	if res {
		cookie := &http.Cookie{
			Name : "username",
			Value: ctx.PostForm("username"),
			Path: "/",
			MaxAge: 300,
			HttpOnly: true,
		}
		http.SetCookie(ctx.Writer, cookie)
		fmt.Println(cookie.Value)
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "login success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "login failed",
		})
	}
}

func Recharge(ctx *gin.Context){
	res := service.Recharge(ctx)
	if res {
		ctx.JSON(http.StatusOK,gin.H{
			"message" : "recharge success",
		})
	} else {
		ctx.JSON(http.StatusOK,gin.H{
			"message" : "recharge failed",
		})
	}
}