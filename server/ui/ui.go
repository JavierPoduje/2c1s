package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
	"github.com/javierpoduje/2c1s/server/logger"
)

type TickMsg time.Time

type Model struct {
	terminalHeight       int
	terminalWidth        int
	logger               *logger.Logger
	running              bool
	sendMessageToClients func()
	actionButtonLabel    string
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
		actionButtonLabel:    "[S]tart",
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

			if m.running {
				m.actionButtonLabel = "[S]top"
			} else {
				m.actionButtonLabel = "[S]tart"
			}

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
	return gloss.Place(
		m.terminalWidth, m.terminalHeight,
		gloss.Center, gloss.Center,
		gloss.JoinVertical(
			gloss.Center,
			titleComp("2 Clients 1 Server"),
			gloss.JoinHorizontal(
				gloss.Center,
				subtitleComp("Conway's Game of Life"),
				gloss.JoinVertical(
					gloss.Left,
					ActionButton(m.actionButtonLabel),
					Dimenssion("Height ", "0"),
					Dimenssion("Width ", "0"),
				),
			),
		),
	)
}
