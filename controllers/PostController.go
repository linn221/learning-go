package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/linn221/go-blog/helpers"
	"github.com/linn221/go-blog/models"
)

func IndexPost(ctx *gin.Context) {
	posts, err := models.GetAllPosts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func GetPost(ctx *gin.Context) {
	post, err := models.GetPostById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

func CreatePost(ctx *gin.Context) {
	var input models.Post
	// binding
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// validating struct
	err = validator.New().Struct(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helpers.ProcessValidationErrors(err)})
		return
	}
	err = input.CreatePost()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, input)
}

func UpdatePost(ctx *gin.Context) {
	var input models.Post
	// binding
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// validating struct
	err = validator.New().Struct(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helpers.ProcessValidationErrors(err)})
		return
	}

	input.ID = helpers.StrToUInt(ctx.Param("id"))
	err = input.UpdatePost()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, input)
}

func DeletePost(ctx *gin.Context) {
	var input models.Post
	input.ID = helpers.StrToUInt(ctx.Param("id"))
	err := input.DeletePost()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(204)
}
