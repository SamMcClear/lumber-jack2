package network

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ReadUserIP(r *http.Request) string {
	UserIP := r.Header.Get("X-Real-Ip")
	if UserIP == "" {
		UserIP = r.Header.Get("X-Forwarded-For")
	}
	if UserIP == "" {
		UserIP = r.RemoteAddr
		fmt.Printf("User IP: %v", UserIP)
	}
	return UserIP
}
func Itraffic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Received request: Method=%s, URL=%s, RemoteAddr=%s", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
		TTC := time.Since(start)

		log.Printf("Request proceed in: %v", TTC)

	})
}

type NetData struct {
	BytesSent     int
	BytesReceived int
	Resport       int
	ReqPort       int
}

func GetNetData() (*NetData, error) {
	return &NetData{
		BytesSent:     100,
		BytesReceived: 200,
		Resport:       8080,
		ReqPort:       8080,
	}, nil
}
