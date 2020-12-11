package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"homework7/models"
)

//将留言和user匹配
func SendMessage(ctx *gin.Context) bool {
	username, err := ctx.Cookie("username")
	if err != nil {
		return false
	}
	fmt.Println(username)
	message := ctx.PostForm("message")
	res := models.SaveMessage(username, message)
	return res
}

//输入用户名查看留言
func ViewMessage(ctx *gin.Context) string {
	username := ctx.PostForm("username")
	messages := models.ViewMessage(username)
	bytes, _ := json.Marshal(messages)
	return string(bytes)
}