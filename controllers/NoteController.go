package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StoreRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}

func CreateNote(context *gin.Context) {
	var request StoreRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.String(200, "hello world")
}
