package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tocken := strings.TrimPrefix(authHeader, "Bearer")
		fmt.Println(tocken)
		next.ServeHTTP(w, r)
	})
}
