package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Method: " + r.Method + " URI: " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}