package conways

type Game struct {
	Board [][]byte
}

func NewGame() *Game {
	return &Game{
		Board: newBoard(),
	}
}

func newBoard() [][]byte {
	board := [][]byte{}

	for i := 0; i < BoardWidth; i++ {
		row := []byte{}
		for j := 0; j < BoardHeight; j++ {
			row = append(row, byte(0))
		}
		board = append(board, row)
	}
	return board
}
