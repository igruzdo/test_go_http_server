package middleware

import (
	"context"
	"http_server/configs"
	"http_server/pakages/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnAuth(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer") {
			writeUnAuth(w)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)

		if !isValid {
			writeUnAuth(w)
			return
		}

		newContext := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(newContext)

		next.ServeHTTP(w, req)
	})
}
