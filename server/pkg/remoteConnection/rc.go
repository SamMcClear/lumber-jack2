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
	RemoteAddress  string
	LocalAddress   string
	RemotePort     int
	LocalPort      int
	RemoteHost     string
	LocalHost      string
	RemoteIP       string
	LocalIP        string
	RemoteProtocol string
	LocalProtocol  string
	RemoteUser     string
	LocalUser      string
}

type SSHAttempt struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
	Username  string `json:"username"`
	SourceIP  string `json:"source_ip"`
	Result    string `json:"result"`
}

func SSHAttemptHandler() ([]SSHAttempt, error) {
	var sshAttempts []SSHAttempt

	rcTrace, err := os.Open("sshAttempts.json")
	if err != nil {
		slog.Error("Error querying database:", "err", err)
		return nil, err
	}
	defer rcTrace.Close()

	return sshAttempts, nil
}

// Exported function for getting IP
func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		ips := strings.Split(ip, ",")
		return strings.TrimSpace(ips[0])
	}

	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		slog.Error("Error parsing RemoteAddr:", "err", err)
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
