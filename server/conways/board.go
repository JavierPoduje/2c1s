package conways

import (
	"fmt"
	"strings"
)

type Board [][]byte

func (b Board) String() string {
	boardAsString := strings.Builder{}
	boardAsString.WriteString("\n")
	for _, row := range b {
		for cellIdx, cell := range row {
			var char string
			if cell == byte(0) {
				char = "0"
			} else {
				char = "1"
			}

			if cellIdx != len(row)-1 {
				char = fmt.Sprintf("%v, ", char)
			}

			boardAsString.WriteString(char)
		}
		boardAsString.WriteString("\n")
	}
	return boardAsString.String()
}

func (b Board) Flatten() []byte {
	message := []byte{}
	for _, row := range b {
		message = append(message, row...)
	}
	return message
}

func (b Board) alive(y, x int) bool {
	return b.get(y, x) == byte(1)
}

func (b Board) get(y, x int) byte {
	return b[y][x]
}

func (b Board) rows() int {
	return len(b)
}

func (b Board) cols() int {
	return len(b[0])
}

func (b Board) Width() int {
	return len(b[0])
}

func (b Board) Height() int {
	return len(b)
}

func newBoard(width, height int) *Board {
	board := Board{}
	seed := seedCoords()

	for y := 0; y < width; y++ {
		row := []byte{}
		for x := 0; x < height; x++ {
			var currSeed []int
			if len(seed) > 0 {
				currSeed = seed[0]
			}

			if len(currSeed) > 0 && currSeed[0] == y && currSeed[1] == x {
				row = append(row, byte(1))
				seed = seed[1:]
			} else {
				row = append(row, byte(0))
			}

		}
		board = append(board, row)
	}

	return &board
}

func seedCoords() [][]int {
	return [][]int{
		{0, 0},
		{1, 1},
		{1, 2},
		{2, 1},
		{2, 2},
		{3, 3},
	}
}

func (b Board) inBounds(y, x int) bool {
	return x >= 0 && x < b.cols() && y >= 0 && y < b.rows()
}
