package ui

import (
	"strconv"

	gloss "github.com/charmbracelet/lipgloss"
)

func titleComp(text string) string {
	return gloss.NewStyle().
		Foreground(foregroundColor()).
		Bold(true).
		Render(text)
}

func subtitleComp(text string) string {
	return gloss.NewStyle().
		Foreground(magentaColor()).
		Render(text)
}

func ActionButton(label string) string {
	return ActionButtonStyles().
		Render(label)
}

func Dimenssion(label string, dimenssion int) string {
	return gloss.JoinHorizontal(
		gloss.Center,
		DimenssionButton(label),
		DimenssionValue(dimenssion),
	)
}

func DimenssionButton(label string) string {
	return ActionButtonStyles().
		Render(label)
}

func DimenssionValue(value int) string {
	valueAsString := strconv.Itoa(value)
	return gloss.NewStyle().
		Foreground(secondaryForegroundColor()).
		Render(valueAsString)
}
