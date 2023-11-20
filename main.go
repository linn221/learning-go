package main

import (
	"notes-app/controllers"
	"notes-app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// create router
	r := gin.Default()
	// connect to database
	utils.ConnectDatabase()

	r.POST("/", controllers.CreateNote)
	r.Run()
}
