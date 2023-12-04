package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linn221/auth/controller"
	"github.com/linn221/auth/model"
)

func main() {
	model.ConnectDB()
	r := gin.Default()
	r.POST("/user", controller.CreateUser)
	r.GET("/user", controller.ListUsers)
	r.POST("/login", controller.AttemptLogin)
	r.Run(":8080")
}
