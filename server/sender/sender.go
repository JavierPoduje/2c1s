package sender

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/javierpoduje/2c1s/server/conways"
)

const TickInterval = (1 * time.Second) / 3

type Server struct {
	clients    []net.Conn
	clientsMtx sync.Mutex
	game       *conways.Game
	width      int
	height     int
}

// message: [width, height, board]
func NewServer(width, height int) *Server {
	return &Server{
		clients:    []net.Conn{},
		clientsMtx: sync.Mutex{},
		game:       conways.NewGame(width, height),
		width:      width,
		height:     height,
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

	ticker := time.NewTicker(TickInterval)
	defer ticker.Stop()

	framNum := 0

	go func() {
		for range ticker.C {
			s.SendMessageToClients(framNum)
			framNum++
		}
	}()

	// accept new clients
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

// Probably, this function should receive width and hegiht as parameters later...
func (s *Server) SendMessageToClients(frameNum int) {
	s.clientsMtx.Lock()
	defer s.clientsMtx.Unlock()

	// don't do anything if there are no clients connected
	if len(s.clients) == 0 {
		return
	}

	// TODO: the width and the height should be updated in the "admin", not here
	if frameNum == 0 {
		// don't do nothing
	} else if frameNum == 15 {
		s.game.Update(s.game.Board.Height()+2, s.game.Board.Width()+2)
	} else {
		s.game.Update(s.game.Board.Height(), s.game.Board.Width())
	}

	for _, conn := range s.clients {
		_, err := conn.Write(s.buildMessage())
		if err != nil {
			fmt.Println("Error writing to client:", err)
			conn.Close()
			s.clients = removeClient(s.clients, conn)
		}
	}
}

func (s *Server) buildMessage() []byte {
	widthAsByte := byte(s.game.Board.Width())
	heightAsByte := byte(s.game.Board.Height())
	flattenBoard := s.game.Board.Flatten()
	return append([]byte{widthAsByte, heightAsByte}, flattenBoard...)
}

func removeClient(clients []net.Conn, clientToRemove net.Conn) []net.Conn {
	for i, conn := range clients {
		if conn == clientToRemove {
			return append(clients[:i], clients[i+1:]...)
		}
	}
	return clients
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Placeholder for any client-specific handling. could be useful later...
	select {}
}
