package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kyawthanttin/retail-pro/backend/models"
)

func GetAllAccounts(context *gin.Context) {

	users, err := models.GetAllAccounts()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success", "data": users})
}

func GetAccount(context *gin.Context) {

	id := context.Param("id")
	if len(strings.TrimSpace(id)) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID is empty or null"})
	}

	model, err := models.GetAccount(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "success", "data": model})
}

func CreateAccount(context *gin.Context) {

	var input models.Account
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := input.CreateAccount()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "create success"})
}

func UpdateAccount(context *gin.Context) {

	var input models.Account
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := input.UpdateAccount()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "update success"})
}

func DeleteAccount(context *gin.Context) {

	var input models.Account
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := input.DeleteAccount()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
