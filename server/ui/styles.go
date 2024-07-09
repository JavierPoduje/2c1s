package ui

import "github.com/charmbracelet/lipgloss"

func ActionButtonStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(foregroundColor()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(cyanColor()).
		Align(lipgloss.Center, lipgloss.Center).
		Width(12).
		Height(3)
}
