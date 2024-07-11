package ui

import gloss "github.com/charmbracelet/lipgloss"

func BoardStyles() gloss.Style {
	return gloss.NewStyle().
		MarginRight(2).
		Align(gloss.Center, gloss.Center)
}

func ActionButtonStyles() gloss.Style {
	return gloss.NewStyle().
		Foreground(foregroundColor()).
		BorderStyle(gloss.RoundedBorder()).
		BorderForeground(cyanColor()).
		Align(gloss.Center, gloss.Center).
		Width(12).
		Height(3)
}
