package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		fmt.Println("Logging")
		next.ServeHTTP(wrapper, r)
		log.Println(time.Since(start))
	})
}
