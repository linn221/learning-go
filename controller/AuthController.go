package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linn221/auth/model"
)

func AttemptLogin(c *gin.Context) {
	var attempt model.User
	if err := c.ShouldBindJSON(&attempt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	token, err := attempt.Login()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}
