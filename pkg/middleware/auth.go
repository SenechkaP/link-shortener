package middleware

import (
	"advpractice/configs"
	"advpractice/pkg/jwt"
	"log"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		log.Println(isValid)
		log.Println(data)
		next.ServeHTTP(w, r)
	})
}
