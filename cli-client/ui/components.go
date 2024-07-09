package ui

import "github.com/charmbracelet/lipgloss"

// the char is the biggest possible ascii square
const CellChar = "██"

func AliveCell() string {
	return cell(AliveCellColor())
}

func DeadCell() string {
	return cell(DeadCellColor())
}

func cell(color lipgloss.Color) string {
	return lipgloss.
		NewStyle().
		Foreground(color).
		Render(CellChar)
}
