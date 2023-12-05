package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linn221/go-blog/models"
)

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

func CurrentUser(c *gin.Context) {
	authId := c.GetUint("auth_id")
	user, err := models.GetUserById(authId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
