package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/javierpoduje/2c1s/server/logger"
)

type TickMsg time.Time

type Model struct {
	terminalHeight       int
	terminalWidth        int
	logger               *logger.Logger
	running              bool
	sendMessageToClients func()
}

func (m Model) tick() tea.Cmd {
	const tickInterval = (1 * time.Second) / 3
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func NewModel(messageToClientsCallback func()) Model {
	return Model{
		terminalHeight:       0,
		terminalWidth:        0,
		logger:               logger.NewLogger("debug.log"),
		running:              false,
		sendMessageToClients: messageToClientsCallback,
	}
}

func (m Model) Init() tea.Cmd {
	return m.tick()
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
		case "s":
			m.running = !m.running
			return m, m.tick()
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case TickMsg:
		if !m.running {
			return m, nil
		}

		m.sendMessageToClients()
		return m, m.tick()
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
