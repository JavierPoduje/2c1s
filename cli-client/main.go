package main

import (
	"fmt"
	"net"
)

const (
	address = "127.0.0.1:12345"
)

func main() {
	StartClient()
}

// StartClient starts the client and listens for messages from the server
func StartClient() {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server at", address)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Println("Received:", string(buffer[:n]))
	}
}
