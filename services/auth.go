package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDE3OTE5MDEsInVzZXJfaWQiOjd9.lp1lnTGyV2FYaU1PLFdWyK8aCO0j1EnhNwVPwllHT0k")

func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24),
	})

	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims, err
}
