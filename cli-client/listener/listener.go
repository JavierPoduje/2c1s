package listener

import (
	"fmt"
	"net"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/javierpoduje/2c1s/cli-client/logger"
	"github.com/javierpoduje/2c1s/cli-client/model"
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
		p := tea.NewProgram(model.NewModel(), tea.WithAltScreen())
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

		c.teaProgram.Send(model.ServerMsg(buffer[:n]))
	}
}
