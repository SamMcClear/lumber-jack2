package network

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func IncomingIP(r *http.Request) string {
	ReqIP := r.Header.Get("X-Real-Ip")
	if ReqIP == "" {
		ReqIP = r.Header.Get("X-Forwarded-For")
	}
	if ReqIP == "" {
		ReqIP = r.RemoteAddr
		fmt.Printf("User IP: %v\n", ReqIP)
	}
	return ReqIP
}

type ReqData struct {
	BytesSent     int
	BytesReceived int
	Resport       int64
	ReqPort       int64
	NumbOfReq     int16
	Method        string
	URL           string
	RemoteAddr    string
	timeStampData time.Time
}

func GetReqData(r *http.Request) (*ReqData, error) {
	// Log the incoming request details
	log.Printf("Received request: Method=%s, URL=%s, RemoteAddr=%s", r.Method, r.URL, r.RemoteAddr)

	// Create and populate the ReqData struct
	logData := &ReqData{
		Method:        r.Method,
		URL:           r.URL.String(),
		RemoteAddr:    r.RemoteAddr,
		timeStampData: time.Now(),
	}

	// Return the populated ReqData struct
	return logData, nil
}

func setupLogging() (*os.File, error) {
	// Open (or create) the log file
	file, err := os.OpenFile("NetworkConn.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	// Set the log output to the file
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // Optional: Includes date, time, and file name/line number in logs

	return file, nil
}
