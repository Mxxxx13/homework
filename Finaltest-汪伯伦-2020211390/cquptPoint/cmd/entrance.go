package cmd

import (
	"github.com/gin-gonic/gin"
	"homework7/controller"
)

func Entrance(){
	router := gin.Default()
	router.POST("/register",controller.Register)
	router.GET("/login",controller.Login)
	router.GET("/recharge",controller.Recharge)
	router.Run(":8080")
}