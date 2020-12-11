package cmd

import (
	"github.com/gin-gonic/gin"
	"homework7/controller"
)

func Entrance(){
	router := gin.Default()
	router.POST("/register",controller.Register)
	router.GET("/login",controller.Login)
	router.POST("/sendMessage", controller.SendMessage)
	router.GET("viewMessage",controller.ViewMessage)
	router.Run(":8080")
}