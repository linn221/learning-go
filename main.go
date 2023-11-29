package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linn221/go-blog/controllers"
	"github.com/linn221/go-blog/models"
)

func main() {
	models.ConnectDB()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello world")
	})
	r.POST("/category", controllers.CreateCategory)
	r.GET("/category", controllers.IndexCategory)
	r.GET("/category/:id", controllers.GetCategory)
	r.PUT("/category/:id", controllers.UpdateCategory)
	r.DELETE("/category/:id", controllers.DeleteCategory)

	r.POST("/post", controllers.CreatePost)
	r.GET("/post", controllers.IndexPost)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)

	r.GET("/db-reset", models.FreshDB)
	r.Run(":8001")
}
