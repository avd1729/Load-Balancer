package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var servers = []string{
	"http://127.0.0.1:5001",
	"http://127.0.0.1:5002",
	"http://127.0.0.1:5003",
}

var currentIndex int = 0

func getServer() string {
	server := servers[currentIndex]
	currentIndex = (currentIndex + 1) % len(servers)
	return server
}

func checkServerHealth(server string) bool {
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(server + "/status")
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	defer resp.Body.Close()

	return true
}

func loadBalance(w http.ResponseWriter, r *http.Request) {
	for range servers {
		curr := getServer()
		if checkServerHealth(curr) {
			client := http.Client{
				Timeout: 3 * time.Second,
			}

			resp, err := client.Get(curr)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			defer resp.Body.Close()

			// Forward response to the client
			w.WriteHeader(resp.StatusCode)
			io.Copy(w, resp.Body)
			return
		}
	}

	http.Error(w, "No healthy backend servers available", http.StatusServiceUnavailable)
}

func main() {
	http.HandleFunc("/", loadBalance)
	fmt.Println("Load balancer running on :8080")
	http.ListenAndServe(":8080", nil)
}
