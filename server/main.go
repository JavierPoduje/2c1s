package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	address = "127.0.0.1:12345"
)

var (
	clients    []net.Conn
	clientsMtx sync.Mutex
)

func main() {
	StartServer()
}

// StartServer starts the server and listens for clients
func StartServer() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started, listening on", address)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			sendToClients()
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New client connected:", conn.RemoteAddr().String())

		clientsMtx.Lock()
		clients = append(clients, conn)
		clientsMtx.Unlock()

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Placeholder for any client-specific handling
	select {}
}

func sendToClients() {
	clientsMtx.Lock()
	defer clientsMtx.Unlock()

	message := []byte("2 clients; 1 server")

	for _, conn := range clients {
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println("Error writing to client:", err)
			conn.Close()
			// Remove the client from the slice
			clients = removeClient(clients, conn)
		}
	}
}

func removeClient(clients []net.Conn, target net.Conn) []net.Conn {
	for i, conn := range clients {
		if conn == target {
			return append(clients[:i], clients[i+1:]...)
		}
	}
	return clients
}
