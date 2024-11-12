package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SamMcClear/lumber-jack2/pkg/cpu"
	"github.com/SamMcClear/lumber-jack2/pkg/network"
	"github.com/SamMcClear/lumber-jack2/pkg/userSpace"
)

type UserSpaceData struct {
	Username string
	LogTS    string
	IP       string
}

// dataXchange
func dataXC(ud *userSpace.UserData) {
	fmt.Printf("User data Xchange: %v", ud)
}

func netReq(n *http.Request) {
	fmt.Printf("incoming request: %v", n)
}

func main() {
	user := UserSpaceData{
		Username: "Sam",
		LogTS:    "2021-09-01T12:00:00Z",
		IP:       "",
	}

	userData, err := userSpace.GetUserData()
	if err != nil {
		fmt.Printf("Error fetching user data: %v\n", err)
		return
	}

	fmt.Print(user)
	dataXC(userData)

	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		cpuInfo := cpu.GetUsage()

		response := map[string]string{
			"message": "Hello from Go!",
			"cpuInfo": cpuInfo,
		}

		json.NewEncoder(w).Encode(response)
	})

	http.Handle("/api/hello", network.Itraffic(helloHandler))

	fmt.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
