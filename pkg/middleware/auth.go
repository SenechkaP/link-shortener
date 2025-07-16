package middleware

import (
	"advpractice/configs"
	"advpractice/pkg/jwt"
	"advpractice/pkg/res"
	"context"
	"net/http"
	"strings"
)

type key string

const ContextEmailKey key = "ContextEmailKey"

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			res.JsonDump(w, "TOKEN IS EMPTY", http.StatusUnauthorized)
			return
		}

		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			res.JsonDump(w, "TOKEN IS NOT VALID", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		newReq := r.WithContext(ctx)
		next.ServeHTTP(w, newReq)
	})
}
