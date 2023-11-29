package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linn221/go-blog/helpers"
	"github.com/linn221/go-blog/models"
)

func CreateCategory(context *gin.Context) {
	var input models.Category
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = input.CreateCategory()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, input)
}

func IndexCategory(context *gin.Context) {
	categories, err := models.GetAllCategories()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, categories)
}

func GetCategory(context *gin.Context) {
	category, err := models.GetCategoryById(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, category)
}

func UpdateCategory(context *gin.Context) {
	var category models.Category
	err := context.ShouldBindJSON(&category)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	category.ID = helpers.StrToUInt(context.Param("id"))
	err = category.UpdateCategory()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, category)
}

func DeleteCategory(context *gin.Context) {
	category := models.Category{
		ID: helpers.StrToUInt(context.Param("id")),
	}
	err := category.DeleteCategory()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.Status(204)
}
