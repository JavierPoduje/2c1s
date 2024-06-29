package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/javierpoduje/2c1s/server/conways"
	"github.com/joho/godotenv"
)

const (
	BoardWidth  = 4
	BoardHeight = 4
)

type Server struct {
	clients         []net.Conn
	clientsMtx      sync.Mutex
	game            *conways.Game
	firstFrameShown bool
}

// message: [width, height, board]
func NewServer() *Server {
	return &Server{
		clients:         []net.Conn{},
		clientsMtx:      sync.Mutex{},
		game:            conways.NewGame(BoardWidth, BoardHeight),
		firstFrameShown: false,
	}
}

func (s *Server) SendMessageToClients() {
	s.clientsMtx.Lock()
	defer s.clientsMtx.Unlock()

	// don't do anything if there are no clients connected
	if len(s.clients) == 0 {
		return
	}

	if s.firstFrameShown {
		s.game.Update(BoardHeight, BoardWidth)
	} else {
		s.firstFrameShown = true
	}

	boardAsMessage := s.game.Board.Flatten()
	message := append([]byte{BoardWidth, BoardHeight}, boardAsMessage...)

	for _, conn := range s.clients {
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println("Error writing to client:", err)
			conn.Close()
			s.clients = removeClient(s.clients, conn)
		}
	}
}

func (s *Server) Start() {
	addr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started, listening on", addr)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			s.SendMessageToClients()
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New client connected:", conn.RemoteAddr().String())

		s.clientsMtx.Lock()
		s.clients = append(s.clients, conn)
		s.clientsMtx.Unlock()

		go handleClient(conn)
	}

}

func (s *Server) BoardToMessage() []byte {
	message := []byte{}
	for _, row := range *s.game.Board {
		message = append(message, row...)
	}
	return message
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := NewServer()
	s.Start()
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Placeholder for any client-specific handling. could be useful later...
	select {}
}

func removeClient(clients []net.Conn, clientToRemove net.Conn) []net.Conn {
	for i, conn := range clients {
		if conn == clientToRemove {
			return append(clients[:i], clients[i+1:]...)
		}
	}
	return clients
}
