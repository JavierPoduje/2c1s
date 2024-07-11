package ui

import (
	"strings"

	gloss "github.com/charmbracelet/lipgloss"
)

// the char is the biggest possible ascii square
const CellChar = "██"

func Board(height, width int, boardSlice []byte) string {
	board := strings.Builder{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := y*width + x
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
	return gloss.Place(
		terminalWidth, terminalHeight,
		gloss.Center, gloss.Center,
		content,
	)
}

func AliveCell() string {
	return cell(AliveCellColor())
}

func DeadCell() string {
	return cell(DeadCellColor())
}

func cell(color gloss.Color) string {
	return gloss.
		NewStyle().
		Foreground(color).
		Render(CellChar)
}
