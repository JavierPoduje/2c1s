package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	terminalHeight int
	terminalWidth  int
}

func NewModel() Model {
	return Model{
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
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return lipgloss.Place(
		m.terminalWidth, m.terminalHeight,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			titleComp("2 Clients 1 Server"),
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				subtitleComp("Conway's Game of Life"),
				ActionButton("[S]tart"),
			),
		),
	)
}
