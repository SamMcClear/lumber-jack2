package network

import (
	"log"
	"net/http"
	"time"
)

func Itraffic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Received request: Method=%s, URL=%s, RemoteAddr=%s", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
		TTC := time.Since(start)

		log.Printf("Request proceed in: %v", TTC)

	})
}
