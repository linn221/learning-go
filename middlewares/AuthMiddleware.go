package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/linn221/go-blog/utils/token"
)

func extractBearer(c *gin.Context) string {
	tokenQuery := c.Query("token")
	if tokenQuery != "" {
		return tokenQuery
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func Auth() gin.HandlerFunc {
	fmt.Println("auth middleware running")
	return func(c *gin.Context) {
		tokenStr := extractBearer(c)
		userid, err := token.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: " + err.Error()})
			c.Abort()
			return
		}
		c.Set("auth_id", userid)
		c.Next()
	}
}
