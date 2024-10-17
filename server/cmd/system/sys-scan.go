package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SamMcClear/lumber-jack2/pkg/cpu"
	"github.com/SamMcClear/lumber-jack2/pkg/network"
)

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(network.Itraffic(next http.HandlerFunc()))
		cpuInfo := cpu.GetUsage()

		response := map[string]string{
			"message": "Hello from Go!",
			"cpuInfo": cpuInfo,
		}

		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("server running on 8080")
	http.ListenAndServe(":8080", nil)
}
