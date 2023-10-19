package middleware

import (
	"fmt"
	"jobs-api/utils"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		elapsed := time.Since(start)

		// Log the request information
		logMessage := fmt.Sprintf("%s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, elapsed)
		utils.GetLogger().Info(logMessage)
	})
}
