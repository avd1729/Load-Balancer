package main

import "fmt"

var servers = []string{
	"http://127.0.0.1:5001",
	"http://127.0.0.1:5002",
	"http://127.0.0.1:5003",
}

var currServer = 0

func server() string {
	currServer = (currServer + 1) % len(servers)
	return servers[currServer]
}

func main() {
	fmt.Println(servers)
}
