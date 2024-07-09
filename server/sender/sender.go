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
	"github.com/javierpoduje/2c1s/server/ui"
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

	// TODO: consider moving this to the bubbletea program
	//go func() {
	//    for range ticker.C {
	//        s.SendMessageToClients()
	//    }
	//}()

	go func() {
		p := tea.NewProgram(ui.NewModel(s.SendMessageToClients), tea.WithAltScreen())

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
	}

}

// Probably, this function should receive width and hegiht as parameters later...
func (s *Server) SendMessageToClients() {
	s.clientsMtx.Lock()
	defer s.clientsMtx.Unlock()

	// don't do anything if there are no clients connected
	if len(s.clients) == 0 {
		return
	}

	s.game.Update(s.game.Board.Height(), s.game.Board.Width())

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

	return append([]byte{
		widthAsByte,
		heightAsByte,
	}, flattenBoard...)
}

func removeClient(clients []net.Conn, clientToRemove net.Conn) []net.Conn {
	for i, conn := range clients {
		if conn == clientToRemove {
			return append(clients[:i], clients[i+1:]...)
		}
	}
	return clients
}
