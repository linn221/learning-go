package utils

import "github.com/golang-jwt/jwt/v5"

const TOKEN_KEY string = "coffee"

func GenerateToken(userid uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"userid":     userid,
		"username":   username,
		"authorized": true,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(TOKEN_KEY))
}
