package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var servers = []string{
	"http://server1:5001",
	"http://server2:5002",
	"http://server3:5003",
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
			// Construct the target URL with the original path and query
			targetURL := curr + r.URL.Path
			if r.URL.RawQuery != "" {
				targetURL += "?" + r.URL.RawQuery
			}

			// Create a new request with the same method and body
			proxyReq, err := http.NewRequest(r.Method, targetURL, r.Body)
			if err != nil {
				fmt.Println("Error creating request:", err)
				continue
			}

			// Copy the original headers
			for name, values := range r.Header {
				for _, value := range values {
					proxyReq.Header.Add(name, value)
				}
			}

			// Forward the request to the backend
			client := http.Client{
				Timeout: 3 * time.Second,
			}

			resp, err := client.Do(proxyReq)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			defer resp.Body.Close()

			// Copy response headers back to client
			for name, values := range resp.Header {
				for _, value := range values {
					w.Header().Add(name, value)
				}
			}

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
