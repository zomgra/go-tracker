package web

import (
	"log"
	"net/http"
	"time"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s - %s (%s) Time beetween excecuting: %v \n", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}
