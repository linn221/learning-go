package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/linn221/go-blog/helpers"
	"github.com/linn221/go-blog/models"
)

func CreateUser(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := validator.New().Struct(input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": helpers.ProcessValidationErrors(err)})
		return
	}

	err := input.CreateUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, input)
}

func GetAllUsers(ctx *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func AttemptLogin(c *gin.Context) {
	form := struct {
		Name     string `validate:"required,min=4" json:"name"`
		Password string `validate:"required,min=8" json:"password"`
	}{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	attemptUser := models.User{Name: form.Name, Password: form.Password}
	tokenStr, err := attemptUser.Login()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   tokenStr,
	})
}
