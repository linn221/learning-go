package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linn221/go-blog/helpers"
	"github.com/linn221/go-blog/models"
)

func CreateCategory(ctx *gin.Context) {
	var input models.Category
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = input.CreateCategory()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, input)
}

func IndexCategory(ctx *gin.Context) {
	categories, err := models.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func GetCategory(ctx *gin.Context) {
	category, err := models.GetCategoryById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func UpdateCategory(ctx *gin.Context) {
	var category models.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	category.ID = helpers.StrToUInt(ctx.Param("id"))
	err = category.UpdateCategory()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func DeleteCategory(ctx *gin.Context) {
	var category models.Category
	category.ID = helpers.StrToUInt(ctx.Param("id"))
	err := category.DeleteCategory()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(204)
}
