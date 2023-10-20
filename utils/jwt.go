package utils

import (
	"jobs-api/api/models"
	"jobs-api/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours from now)
	}

	secretKey := []byte(config.GetConfig().App.JWTKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
