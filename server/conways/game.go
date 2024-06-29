package conways

import "log"

type Game struct {
	Board [][]byte
}

func NewGame(width, height int) *Game {
	return &Game{
		Board: newBoard(width, height),
	}
}

func (g *Game) Update(height, width int) {
	newBoard := [][]byte{}
	for y := 0; y < width; y++ {
		row := []byte{}
		for x := 0; x < height; x++ {
			row = append(row, g.getNextCellState(x, y))
		}
		newBoard = append(newBoard, row)
	}
	g.Board = newBoard
}

// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// Any live cell with two or three live neighbours lives on to the next generation.
// Any live cell with more than three live neighbours dies, as if by overpopulation.
// [not implemented. I don't think its necessary but let's keep an eye on it...]
// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func (g *Game) getNextCellState(y, x int) byte {
	neighbours := g.getNeighbours(y, x)

	if neighbours < 2 {
		return byte(0)
	} else if neighbours == 2 || neighbours == 3 {
		return byte(1)
	} else if neighbours > 3 {
		return byte(0)
	}

	log.Panic("Unreachable")

	return byte(0)
}

func (g *Game) getNeighbours(y, x int) int {
	directions := [][]int{
		{-1, 0},  // up
		{-1, 1},  // up-right
		{0, 1},   // right
		{1, 1},   // bottom-right
		{1, 0},   // bottom
		{1, -1},  // bottom-left
		{0, -1},  // left
		{-1, -1}, // up-left
	}

	neighbours := 0
	for _, dir := range directions {
		neighbourY := y + dir[0]
		neighbourx := x + dir[1]

		if g.inBounds(neighbourx, neighbourY) && g.Board[neighbourY][neighbourx] == byte(1) {
			neighbours++
		}
	}
	return neighbours
}

func (g *Game) inBounds(x, y int) bool {
	return x >= 0 && x < len(g.Board[0]) && y >= 0 && y < len(g.Board)
}

func newBoard(width, height int) [][]byte {
	board := [][]byte{}
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
	return board
}

func seedCoords() [][]int {
	return [][]int{
		{4, 4},
		{4, 5},
		{5, 3},
		{5, 6},
		{6, 3},
		{6, 6},
		{7, 3},
		{7, 4},
		{7, 5},
	}
}
