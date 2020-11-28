package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户的信息
type uesr struct {
	Name     string `form:"用户名" json:"user" binding:"required"`
	Password string `form:"密码" json:"password" binding:"required"`
}

//选择登录还是注册
type selection struct {
	Selection string `form:"选项" json:"selection" binding:"required"`
}

type usermessage map[string]string
var um usermessage = make(usermessage)
var u uesr
var s selection

//登录是否成功
func SignIn()bool{
	var b bool
	if um[u.Name] == u.Password{
		b = true
	}else {
		b = false
	}
	return b
}

//检查用户是否已经存在
func IsExist()bool{
	var b bool
	for k,_:= range um {
		if u.Name == k {
			b = true
		}
	}
	return b
}

func main() {
	router := gin.Default()
	router.GET("/test", func(ctx *gin.Context) {
		if err := ctx.ShouldBind(&s); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"选项":s.Selection,
			})
		}

		if err := ctx.ShouldBind(&u); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"用户名": u.Name,
				"密码":  u.Password,
			})
		}
		if s.Selection == "注册"{
			bool:=IsExist()
			if bool == true{
				ctx.JSON(200,"用户已存在")
			}else{
				um[u.Name] = u.Password
				ctx.JSON(200,"注册成功")
			}
		}
		if s.Selection == "登录"{
			if SignIn() == true{
				ctx.JSON(200,"登录成功")
			}else if SignIn() == false{
				ctx.JSON(200,"用户名或密码错误")
			}
		}
	})
	router.Run(":8080")
}
