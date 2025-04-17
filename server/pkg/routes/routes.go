package routes

import (
	"encoding/json"
	"net/http"

	"github.com/SamMcClear/lumber-jack2/pkg/cpu"
	"github.com/SamMcClear/lumber-jack2/pkg/network"
	"github.com/SamMcClear/lumber-jack2/pkg/userSpace"
)

// RegisterRoutes sets up all the API routes
func RegisterRoutes() {
	// Home endpoint
	http.HandleFunc("/api/home", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Fetch real-time data
		cpuInfo := cpu.GetUsage()
		userData, _ := userSpace.GetUserData()
		userIP := network.ReadUserIP(r)

		// Construct the response
		response := map[string]string{
			"message":  "Welcome to LumberJack!", // Dynamic message
			"cpuInfo":  cpuInfo,                  // Real-time CPU info
			"UserIP":   userIP,                   // Real-time user IP
			"userInfo": userData.Username,        // Real-time user info
		}

		// Send the response
		json.NewEncoder(w).Encode(response)
	})

	// Route for IP scanning
	http.HandleFunc("/api/ip", func(w http.ResponseWriter, r *http.Request) {
		ip := network.ReadUserIP(r)
		w.Write([]byte("Client IP: \n" + ip))
	})

	// Route for CPU usage
	http.HandleFunc("/api/cpu", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		cpuInfo := cpu.GetUsage()
		json.NewEncoder(w).Encode(map[string]string{
			"cpuInfo": cpuInfo,
		})
	})

	// Route for network usage
	http.HandleFunc("/api/network", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"networkUsage": "", // Replace with dynamic data if available
		})
	})

	// Route for system info
	http.HandleFunc("/api/system", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"systemInfo": "", // Replace with dynamic data
		}
		json.NewEncoder(w).Encode(response)
	})

	// Route for storage info
	http.HandleFunc("/api/storage", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"storageInfo": "", // Replace with dynamic data
		}
		json.NewEncoder(w).Encode(response)
	})

	// Route for user and server info
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		userData, err := userSpace.GetUserData()
		if err != nil {
			http.Error(w, "Failed to fetch user data", http.StatusInternalServerError)
			return
		}
		response := map[string]string{
			"userInfo": userData.Username, // Replace with dynamic data
		}
		json.NewEncoder(w).Encode(response)
	})

	// Example hello route
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		cpuInfo := cpu.GetUsage()
		userData, _ := userSpace.GetUserData()

		response := map[string]string{
			"message":    "Hello from Go!",
			"cpuInfo":    cpuInfo,
			"UserIP":     network.ReadUserIP(r),
			"username":   userData.Username,
			"serverName": "LumberJack Server",
		}

		json.NewEncoder(w).Encode(response)
	})
}
