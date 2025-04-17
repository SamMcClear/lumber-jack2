package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SamMcClear/lumber-jack2/pkg/logging"
	"github.com/SamMcClear/lumber-jack2/pkg/routes"
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

	// Register all routes
	routes.RegisterRoutes()

	// Start the server
	fmt.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		logging.NewErrorLog().WriteLog(err)
		return
	}
}
