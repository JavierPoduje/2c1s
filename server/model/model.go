package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	height int
	width  int
}

func NewModel() Model {
	return Model{
		height: 0,
		width:  0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) HandleWindowResize(msg tea.WindowSizeMsg) {
	m.width, m.height = msg.Width, msg.Height
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
	return "this is bubbletea, baby!"
}
