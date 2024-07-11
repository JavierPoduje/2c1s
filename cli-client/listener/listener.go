package listener

import (
	"fmt"
	"net"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/javierpoduje/2c1s/cli-client/logger"
	"github.com/javierpoduje/2c1s/cli-client/ui"
)

type Listener struct {
	addr       string
	teaProgram *tea.Program
	logger     *logger.Logger
}

func NewClient(addr string) *Listener {
	return &Listener{
		addr:       addr,
		teaProgram: nil,
		logger:     logger.NewLogger("debug.log"),
	}
}

func (c *Listener) Start() {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	go func() {
		p := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())
		c.teaProgram = p

		if _, err := p.Run(); err != nil {
			fmt.Printf("Error starting program: %v", err)
			os.Exit(1)
		}
	}()

	c.logger.Log(fmt.Sprintf("Connected to server at %v", c.addr))

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}

		parsedMessage := c.parseMessage(buffer[:n])
		c.teaProgram.Send(ui.ServerMsg(parsedMessage))
	}
}

func (l Listener) parseMessage(msg []byte) ui.ServerMsg {
	width := int(msg[0])
	height := int(msg[1])
	board := msg[2:]

	serverMessage := ui.ServerMsg{
		Width:  width,
		Height: height,
		Board:  board,
	}

	l.logger.Log(fmt.Sprintf("message: %v\n", serverMessage))

	return serverMessage
}
