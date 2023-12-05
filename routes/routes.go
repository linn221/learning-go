package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/linn221/go-blog/controllers"
	"github.com/linn221/go-blog/middlewares"
)

func StartServer() {
	r := gin.Default()
	protected := r.Group("")
	protected.Use(middlewares.Auth())
	public := r.Group("")

	protected.POST("/user", controllers.CreateUser)
	protected.GET("/user", controllers.GetAllUsers)
	protected.GET("/me", controllers.CurrentUser)

	protected.POST("/category", controllers.CreateCategory)
	protected.GET("/category", controllers.IndexCategory)
	protected.GET("/category/:id", controllers.GetCategory)
	protected.PUT("/category/:id", controllers.UpdateCategory)
	protected.DELETE("/category/:id", controllers.DeleteCategory)

	protected.POST("/post", controllers.CreatePost)
	protected.GET("/post", controllers.IndexPost)
	protected.GET("/post-slug/:slug", controllers.GetPostBySlug)
	protected.GET("/post/:id", controllers.GetPost)
	protected.PUT("/post/:id", controllers.UpdatePost)
	protected.DELETE("/post/:id", controllers.DeletePost)

	protected.POST("/tag", controllers.CreateTag)
	protected.GET("/tag", controllers.IndexTag)
	protected.GET("/tag/:id", controllers.GetTag)
	protected.PUT("/tag/:id", controllers.UpdateTag)
	protected.DELETE("/tag/:id", controllers.DeleteTag)

	public.POST("/login", controllers.AttemptLogin)
	public.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello world")
	})

	r.Run(":8001")

}
