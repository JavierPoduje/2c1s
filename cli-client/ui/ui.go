package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	msgFromServer  ServerMsg
	terminalHeight int
	terminalWidth  int
}

type ServerMsg struct {
	Height int
	Width  int
	Board  []byte
}

func NewModel() Model {
	return Model{
		msgFromServer:  ServerMsg{},
		terminalHeight: 0,
		terminalWidth:  0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) HandleWindowResize(msg tea.WindowSizeMsg) {
	m.terminalWidth, m.terminalHeight = msg.Width, msg.Height
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.HandleWindowResize(msg)
	case ServerMsg:
		m.msgFromServer = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	width := m.msgFromServer.Width
	height := m.msgFromServer.Height
	rawBoard := m.msgFromServer.Board

	return Layout(
		m.terminalWidth,
		m.terminalHeight,
		Board(height, width, rawBoard),
	)
}
