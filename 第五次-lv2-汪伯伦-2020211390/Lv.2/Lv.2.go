package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type uesr struct {
	Name     string `form:"用户名" json:"user" binding:"required"`
	Password string `form:"密码" json:"password" binding:"required"`
}

type selection struct {
	Selection string `form:"选项" json:"selection" binding:"required"`
	Remerber string `form:"记住密码" json:"selection" binding:"required"`
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

//是否记住密码
func IsRemerber()bool{
	var b bool
	if s.Remerber == "是"{
		b = true
	}else if s.Remerber == "否"{
		b = false
	}
	return b
}

func main() {
	router := gin.Default()
	router.GET("/login", func(ctx *gin.Context) {
		 ctx.ShouldBind(&s)
		 ctx.ShouldBind(&u)
		 ctx.JSON(http.StatusOK,gin.H{
			 "code":200,
			 "message":"游客你好！",
		 })
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
			if IsRemerber() == true{
				cookie1 := &http.Cookie{
					Name : "username",
					Value: u.Name,
					Path: "/",
					HttpOnly: true,
				}
				cookie2 := &http.Cookie{
					Name : "password",
					Value: u.Password,
					Path: "/",
					HttpOnly: true,
				}
				http.SetCookie(ctx.Writer, cookie1)
				http.SetCookie(ctx.Writer, cookie2)
				u.Name = cookie1.Value
				u.Password = cookie2.Value
			}

			if SignIn() == true{
				ctx.Request.URL.Path = "/uesr"
				router.HandleContext(ctx)
			}else if SignIn() == false{
				ctx.JSON(200,"用户名或密码错误")
			}
		}
	})
	router.GET("/uesr", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
				"code":200,
				"message":u.Name + "你好！",
		})
	})
	router.Run(":8080")
}

