package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linn221/go-blog/controllers"
	"github.com/linn221/go-blog/models"
)

func main() {
	models.ConnectDB()

	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(200, "hello world")
	})
	r.POST("/category", controllers.CreateCategory)
	r.GET("/category", controllers.IndexCategory)
	r.GET("/category/:id", controllers.GetCategory)
	r.PUT("/category/:id", controllers.UpdateCategory)
	r.DELETE("/category/:id", controllers.DeleteCategory)
	r.Run(":8001")
}
