package middleware

import (
	"log"
	"net/http"
	"strings"
)

func Bearer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		log.Println(token)
		next.ServeHTTP(w, r)
	})
}
