package ui

import "github.com/charmbracelet/lipgloss"

func titleComp(text string) string {
	return lipgloss.NewStyle().
		Foreground(foregroundColor()).
		Bold(true).
		Render(text)
}

func subtitleComp(text string) string {
	return lipgloss.NewStyle().
		Foreground(magentaColor()).
		Render(text)
}

func ActionButton(label string) string {
	return ActionButtonStyles().
		Render(label)
}
