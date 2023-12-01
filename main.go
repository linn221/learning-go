package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/linn221/go-blog/controllers"
	"github.com/linn221/go-blog/models"
	"github.com/linn221/go-blog/seeders"
)

func startServer() {
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
	r.GET("/post-slug/:slug", controllers.GetPostBySlug)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)

	r.POST("/tag", controllers.CreateTag)
	r.GET("/tag", controllers.IndexTag)
	r.GET("/tag/:id", controllers.GetTag)
	r.PUT("/tag/:id", controllers.UpdateTag)
	r.DELETE("/tag/:id", controllers.DeleteTag)

	r.Run(":8001")

}

func main() {
	models.ConnectDB()
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "migrate":
			models.FreshDB()
		case "seed":
			seeders.Run()
		case "migrate:seed":
			models.FreshDB()
			seeders.Run()
		}
		return
	}
	startServer()
}
