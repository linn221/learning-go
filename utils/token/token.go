package token

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

const APP_KEY string = "coffee"

func GenerateAuthToken(userId uint, userName string, expiredDate uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["userid"] = userId
	claims["username"] = userName
	claims["exp"] = expiredDate
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(APP_KEY), nil
	})
	if err != nil {
		return errors.New("parsing error")
	}
	return nil
}
