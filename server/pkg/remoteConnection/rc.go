package remoteConnection

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
)

// RemoteConnection is a struct that holds the connection information

var err error

type RemoteConnection struct {
	// The remote address of the connection
	RemoteAddress string
	// The local address of the connection
	LocalAddress string
	// The remote port of the connection
	RemotePort int
	// The local port of the connection
	LocalPort int
	// The remote host of the connection
	RemoteHost string
	// The local host of the connection
	LocalHost string
	// The remote IP address of the connection
	RemoteIP string
	// The local IP address of the connection
	LocalIP string
	// The remote protocol of the connection
	RemoteProtocol string
	// The local protocol of the connection
	LocalProtocol string
	// The remote user of the connection
	RemoteUser string
	// The local user of the connection
	LocalUser string
}

func getIP(r *http.Request) string {
	// Check if the request contains an X-Forwarded-For header (useful when behind a proxy)
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// X-Forwarded-For can contain multiple IPs, the first one is the client IP
		ips := strings.Split(ip, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check if the request contains an X-Real-IP header (another proxy header)
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	// Get the IP from RemoteAddr (fallback)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		slog.Error("Error parsing RemoteAddr:", "err", "%v")
		return "unknown"
	}

	return ip
}

// NewRemoteConnection creates a new RemoteConnection struct
func NewRemoteConnection() *RemoteConnection {

	// Get the remote address from the environment variable
	remoteAddress := os.Getenv("REMOTE_ADDRESS")

	fmt.Println(remoteAddress)
	return &RemoteConnection{}
}

// GetRemoteConnection returns the remote connection information
func (rc *RemoteConnection) GetRemoteConnection() *RemoteConnection {
	// Get the remote address from the environment variable
	if rc.RemoteAddress == "" {
		slog.Error("Remote address is not found")
	}
	return rc
}

// GetRemoteConnectionInfo returns the remote connection information
func (rc *RemoteConnection) GetRemoteConnectionInfo() *RemoteConnection {
	return rc
}
