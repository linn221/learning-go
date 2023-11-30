package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linn221/go-blog/helpers"
	"github.com/linn221/go-blog/models"
)

func CreateTag(ctx *gin.Context) {
	var input models.Tag
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := input.CreateTag()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, input)
}

func UpdateTag(ctx *gin.Context) {
	var input models.Tag
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	input.ID = helpers.StrToUInt(ctx.Param("id"))
	err := input.UpdateTag()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, input)
}

func DeleteTag(ctx *gin.Context) {
	var input models.Tag
	input.ID = helpers.StrToUInt(ctx.Param("id"))
	err := input.DeleteTag()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(204)
}

func IndexTag(ctx *gin.Context) {
	tags, err := models.GetAllTags()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tags)
}

func GetTag(ctx *gin.Context) {
	tag, err := models.GetTagById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tag)
}
