package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SamMcClear/lumber-jack2/pkg/cpu"
	"github.com/SamMcClear/lumber-jack2/pkg/logging"
	"github.com/SamMcClear/lumber-jack2/pkg/network"
	"github.com/SamMcClear/lumber-jack2/pkg/remoteConnection"
	"github.com/SamMcClear/lumber-jack2/pkg/userSpace"
)

type UserSpaceData struct {
	Username string
	LogTS    string
	IP       string
}

// dataXchange
func dataXC(ud *userSpace.UserData) {
	fmt.Printf("Logged in User: %v\n", ud)
}

func netReq(n *http.Request) {
	fmt.Printf("incoming request: %v\n", n)
}

func netResp(n *http.Response) {
	fmt.Printf("outgoing response: %v\n", n)
}

func ScanIP(r *http.Request) {
	ip := remoteConnection.GetIP(r)
	fmt.Println("Client IP:", ip)
}

func main() {
	user := UserSpaceData{
		Username: "Sam",
		LogTS:    time.Now().String(),
		IP:       "",
	}

	userData, err := userSpace.GetUserData()
	if err != nil {
		fmt.Printf("Error fetching user data: %v\n", err)
		return
	}

	fmt.Println(user)
	dataXC(userData)

	// API route for IP scanning
	http.HandleFunc("/api/ip", func(w http.ResponseWriter, r *http.Request) {
		ScanIP(r)
		w.Write([]byte("Check logs for IP"))
	})

	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		cpuInfo := cpu.GetUsage()

		response := map[string]string{
			"message": "Hello from Go!",
			"cpuInfo": cpuInfo,
			"UserIP":  network.ReadUserIP(r),
		}

		json.NewEncoder(w).Encode(response)
	})

	http.Handle("/api/hello", network.Itraffic(helloHandler))

	fmt.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		logging.NewErrorLog().WriteLog(err)
		return
	}
}
