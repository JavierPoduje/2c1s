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
	boardHeight          int
	boardWidth           int
	seed                 [][]int
	seedWasUpdated       bool
	togglerCoord         []int
	logger               *logger.Logger
	running              bool
	sendMessageToClients func(height, width int, seed [][]int)
	actionButtonLabel    string
}

func (m Model) tick() tea.Cmd {
	const tickInterval = (1 * time.Second) / 3
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func NewModel(messageToClientsCallback func(height, width int, seed [][]int), initialBoardHeight, initialBoardWidth int, seed [][]int) Model {
	return Model{
		terminalHeight:       0,
		terminalWidth:        0,
		seed:                 seed,
		seedWasUpdated:       false,
		boardHeight:          initialBoardHeight,
		boardWidth:           initialBoardWidth,
		togglerCoord:         []int{0, 0},
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
		case "shift+up":
			m.boardHeight++
			m.seed = updateSeedDimensions(m.seed, m.boardHeight, m.boardWidth)
			m.seedWasUpdated = true
			m.togglerCoord = []int{0, 0}
		case "shift+down":
			m.boardHeight--
			m.seed = updateSeedDimensions(m.seed, m.boardHeight, m.boardWidth)
			m.seedWasUpdated = true
			m.togglerCoord = []int{0, 0}
		case "shift+left":
			m.boardWidth--
			m.seed = updateSeedDimensions(m.seed, m.boardHeight, m.boardWidth)
			m.seedWasUpdated = true
			m.togglerCoord = []int{0, 0}
		case "shift+right":
			m.boardWidth++
			m.seed = updateSeedDimensions(m.seed, m.boardHeight, m.boardWidth)
			m.seedWasUpdated = true
			m.togglerCoord = []int{0, 0}
		case " ":
			togglerY, togglerX := m.togglerCoord[0], m.togglerCoord[1]
			seedIsAlive := m.seed[togglerY][togglerX] == 1
			if seedIsAlive {
				m.seed[togglerY][togglerX] = 0
			} else {
				m.seed[togglerY][togglerX] = 1
			}
			m.seedWasUpdated = true
		case "up":
			if m.togglerCoord[0] > 0 {
				m.togglerCoord[0]--
			}
		case "down":
			if m.togglerCoord[0] < m.boardHeight-1 {
				m.togglerCoord[0]++
			}
		case "left":
			if m.togglerCoord[1] > 0 {
				m.togglerCoord[1]--
			}
		case "right":
			if m.togglerCoord[1] < m.boardWidth-1 {
				m.togglerCoord[1]++
			}
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

		if m.seedWasUpdated {
			m.sendMessageToClients(m.boardHeight, m.boardWidth, m.seed)
		} else {
			m.sendMessageToClients(m.boardHeight, m.boardWidth, [][]int{})
		}

		m.seedWasUpdated = false

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
				BoardSeed(m.boardHeight, m.boardWidth, m.seed, m.togglerCoord),
				gloss.JoinVertical(
					gloss.Left,
					ActionButton(m.actionButtonLabel),
					Dimenssion("Height ", m.boardHeight),
					Dimenssion("Width ", m.boardWidth),
				),
			),
		),
	)
}

func updateSeedDimensions(currentSeed [][]int, newHeight, newWidth int) [][]int {
	newSeed := blankSeed(newHeight, newWidth)

	currentHeight := len(currentSeed)
	currentWidth := len(currentSeed[0])

	for y := 0; y < currentHeight; y++ {
		if y >= newHeight {
			break
		}
		for x := 0; x < currentWidth; x++ {
			if x >= newWidth {
				continue
			}

			// take the value of the current board
			value := currentSeed[y][x]
			newSeed[y][x] = value
		}
	}

	return newSeed
}

func blankSeed(height, width int) [][]int {
	seed := make([][]int, height)
	for y := 0; y < height; y++ {
		row := make([]int, width)
		for x := 0; x < width; x++ {
			row[x] = 0
		}
		seed[y] = row
	}
	return seed
}
