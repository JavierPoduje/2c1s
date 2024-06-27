package conways

type Game struct {
	Board [][]byte
}

func NewGame(width, height int) *Game {
	return &Game{
		Board: newBoard(width, height),
	}
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
