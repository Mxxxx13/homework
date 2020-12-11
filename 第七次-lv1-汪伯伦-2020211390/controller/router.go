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
			"code" : 100,
			"message" : "register success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 200,
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
			"code" : 100,
			"message" : "login success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"message" : "login failed",
		})
	}
}

//返回留言后的信息
func SendMessage(ctx *gin.Context){
	res := service.SendMessage(ctx)
	if res {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 100,
			"message" : "send message success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"message" : "send message failed",
		})
	}
}

//返回留言
func ViewMessage(ctx *gin.Context){
	messages := service.ViewMessage(ctx)
	fmt.Println(messages)
	if messages != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     100,
			"messages": messages,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": ctx.PostForm("username") + "未留言过",
		})
	}
}