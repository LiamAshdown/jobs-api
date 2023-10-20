package middleware

import (
	"context"
	"fmt"
	"jobs-api/api/models"
	"jobs-api/config"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		// Strip the Bearer prefix from the token string
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte(config.GetConfig().App.JWTKey), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			userMap := claims["user"].(map[string]interface{})

			user := models.User{
				ID:        int(userMap["id"].(float64)),
				Email:     userMap["email"].(string),
				Password:  userMap["password"].(string),
				CreatedAt: userMap["created_at"].(string),
				UpdatedAt: userMap["updated_at"].(string),
			}

			// Add the user object to the request context
			ctx := context.WithValue(r.Context(), "user", user)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
