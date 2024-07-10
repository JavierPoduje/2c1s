package ui

import gloss "github.com/charmbracelet/lipgloss"

func ActionButtonStyles() gloss.Style {
	return gloss.NewStyle().
		Foreground(foregroundColor()).
		BorderStyle(gloss.RoundedBorder()).
		BorderForeground(cyanColor()).
		Align(gloss.Center, gloss.Center).
		Width(12).
		Height(3)
}
