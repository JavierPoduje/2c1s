package sender

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/javierpoduje/2c1s/cli-client/logger"
	"github.com/javierpoduje/2c1s/server/conways"
	"github.com/javierpoduje/2c1s/server/model"
)

const TickInterval = (1 * time.Second) / 3

type Server struct {
	clients    []net.Conn
	clientsMtx sync.Mutex
	game       *conways.Game
	width      int
	height     int
	logger     *logger.Logger
}

// message: [width, height, board]
func NewServer(width, height int) *Server {
	return &Server{
		clients:    []net.Conn{},
		clientsMtx: sync.Mutex{},
		game:       conways.NewGame(width, height),
		width:      width,
		height:     height,
		logger:     logger.NewLogger("debug.log"),
	}
}

func (s *Server) Start() {
	addr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		s.logger.Log(fmt.Sprintf("Error starting server: %v", err))
		return
	}
	defer listener.Close()

	s.logger.Log(fmt.Sprintf("Server started, listening on %v", addr))

	ticker := time.NewTicker(TickInterval)
	defer ticker.Stop()

	framNum := 0

	// TODO: consider moving this to the bubbletea program
	go func() {
		for range ticker.C {
			s.SendMessageToClients(framNum)
			framNum++
		}
	}()

	go func() {
		p := tea.NewProgram(model.NewModel(), tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			s.logger.Log(fmt.Sprintf("Error starting program: %v", err))
			os.Exit(1)
		}
	}()

	// accept new clients
	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Log(fmt.Sprintf("Error accepting connection: %s", err))
			continue
		}

		s.logger.Log(fmt.Sprintf("New client connected: %s", conn.RemoteAddr().String()))

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
			s.logger.Log(fmt.Sprintf("Error writing to client: %v", err))
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
