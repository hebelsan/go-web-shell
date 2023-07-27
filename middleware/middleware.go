package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(fmt.Sprintf("%s %s %v", r.Method, r.URL.Path, time.Since(start)))
	})
}

func TokenCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("token")
		if header != os.Getenv("SECRET_TOKEN") {
			http.Error(w, "That token is invalid.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
