package token

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

const APP_KEY string = "coffee"

func GenerateAuthToken(userId uint, userName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userid"] = userId
	claims["username"] = userName
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("bearer %v", tokenStr), nil
}

func ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(APP_KEY), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if token.Valid && ok {
		userId := uint(claims["userid"].(float64))
		return userId, nil
	}
	return 0, errors.New("invalid token")
}
