package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// the char is the biggest possible ascii square
const CellChar = "██"

func Board(height, width int, boardSlice []byte) string {
	board := strings.Builder{}
	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			idx := y*height + x
			if boardSlice[idx] == byte(1) {
				board.WriteString(AliveCell())
			} else {
				board.WriteString(DeadCell())
			}
		}
		board.WriteString("\n")
	}
	return board.String()
}

func Layout(terminalWidth, terminalHeight int, content string) string {
	return lipgloss.Place(
		terminalWidth, terminalHeight,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}

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
