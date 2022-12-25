package services

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("my-super-duper-secret-key") // move to env

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = "username" //update username
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}
