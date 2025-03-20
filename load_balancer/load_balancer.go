package main

import (
	"fmt"
	"net/http"
	"time"
)

var servers = []string{
	"http://127.0.0.1:5001",
	"http://127.0.0.1:5002",
	"http://127.0.0.1:5003",
}

var currServer = 0

func getServer() string {
	currServer = (currServer + 1) % len(servers)
	return servers[currServer]
}

func checkServerHealth(server string) bool {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(server)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func main() {
	fmt.Println(servers)
}
