package conways

type Game struct {
	Board *Board
}

func NewGame(height, width int) *Game {
	return &Game{
		Board: newBoard(height, width),
	}
}

func (g *Game) Update(height, width int) {
	newBoard := Board{}

	// update if the board has changed
	if g.Board.Height() != height || g.Board.Width() != width {
		g.Board = g.Board.UpdateBoardDimensions(height, width)
	}

	for y := 0; y < height; y++ {
		row := []byte{}
		for x := 0; x < width; x++ {
			row = append(row, g.getNextCellState(y, x))
		}
		newBoard = append(newBoard, row)
	}
	g.Board = &newBoard
}

// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// Any live cell with two or three live neighbours lives on to the next generation.
// Any live cell with more than three live neighbours dies, as if by overpopulation.
// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func (g Game) getNextCellState(y, x int) byte {
	neighbours := g.getNeighbours(y, x)

	if g.Board.alive(y, x) && neighbours < 2 {
		return byte(0)
	} else if g.Board.alive(y, x) && neighbours == 2 || neighbours == 3 {
		return byte(1)
	} else if g.Board.alive(y, x) && neighbours > 3 {
		return byte(0)
	} else if !g.Board.alive(y, x) && neighbours == 3 {
		return byte(1)
	}

	return byte(0)
}

func (g Game) getNeighbours(y, x int) int {
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
		neighbourX := x + dir[1]

		if g.Board.inBounds(neighbourY, neighbourX) && g.Board.alive(neighbourY, neighbourX) {
			neighbours++
		}
	}
	return neighbours
}
