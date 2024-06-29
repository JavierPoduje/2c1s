package model

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	msg ServerMsg
}

type ServerMsg []byte

func NewModel() Model {
	return Model{
		msg: ServerMsg{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ServerMsg:
		m.msg = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	if len(m.msg) == 0 {
		msg := fmt.Sprintf("m.msg: %v\n", m.msg)
		return msg
	}

	width := int(m.msg[0])
	height := int(m.msg[1])
	board := m.msg[2:]

	str := strings.Builder{}

	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			if board[y*height+x] == byte(1) {
				str.WriteString("🟩")
			} else {
				str.WriteString("⬛")
			}
		}
		str.WriteString("\n")
	}

	return str.String()
}
