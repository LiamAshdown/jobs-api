package utils

import (
	"jobs-api/api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key") // Replace with your own secret key

func GenerateJWTToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours from now)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
