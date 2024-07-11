package ui

import (
	"strconv"
	"strings"

	gloss "github.com/charmbracelet/lipgloss"
)

const CellChar = "██"

func BoardSeed(height, width int, seed [][]int, togglerCoord []int) string {
	boardStr := strings.Builder{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if togglerCoord[0] == y && togglerCoord[1] == x {
				if seed[y][x] == 1 {
					boardStr.WriteString(TogglerAliveCell())
				} else {
					boardStr.WriteString(TogglerDeadCell())
				}
				continue
			}

			if seed[y][x] == 1 {
				boardStr.WriteString(AliveCell())
			} else {
				boardStr.WriteString(DeadCell())
			}
		}
		boardStr.WriteString("\n")
	}

	return BoardStyles().Render(boardStr.String())
}

func TogglerAliveCell() string {
	return cell(TogglerAliveCellColor())
}

func TogglerDeadCell() string {
	return cell(TogglerDeadCellColor())
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

func titleComp(text string) string {
	return gloss.NewStyle().
		Foreground(foregroundColor()).
		Bold(true).
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
