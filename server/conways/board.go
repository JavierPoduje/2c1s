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

func newBoard(height, width int) *Board {
	board := blankBoard(height, width)
	seed := diamondRing(height, width)
	placeSeed(&board, seed)
	return &board
}

func (b *Board) UpdateBoardDimensions(height, width int) *Board {
	newBoard := blankBoard(height, width)

	boardHeight := b.Height()
	boardWidth := b.Width()

	for y := 0; y < height; y++ {
		if y >= boardHeight {
			break
		}

		for x := 0; x < width; x++ {
			if x >= boardWidth {
				newBoard[y][x] = byte(0)
				continue
			}

			// take the value of the current board
			value := (*b)[y][x]
			newBoard[y][x] = value
		}
	}

	return &newBoard
}

func blankBoard(height, width int) Board {
	board := Board{}
	for y := 0; y < height; y++ {
		row := []byte{}
		for x := 0; x < width; x++ {
			row = append(row, byte(0))
		}
		board = append(board, row)
	}
	return board
}

func placeSeed(b *Board, seed [][]int) {
	for y, row := range seed {
		for x, cell := range row {
			if cell == 1 {
				(*b)[y][x] = byte(1)
			}
		}
	}
}

func (b Board) inBounds(y, x int) bool {
	return x >= 0 && x < b.cols() && y >= 0 && y < b.rows()
}
