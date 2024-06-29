package conways

import (
	"reflect"
	"testing"
)

type testContext struct {
	game *Game
}

func (c *testContext) beforeEach() {
	game := Game{
		Board: [][]byte{
			{0, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 0, 1},
			{0, 1, 0, 0, 1, 0},
			{0, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0},
		},
	}
	c.game = &game
}

func testCase(test func(t *testing.T, c *testContext)) func(*testing.T) {
	return func(t *testing.T) {
		context := &testContext{}
		context.beforeEach()
		test(t, context)
	}
}

func TestGame_getNeighbours(t *testing.T) {
	t.Run("corner is read without erroing", testCase(func(t *testing.T, c *testContext) {
		game := c.game
		neighbours := game.getNeighbours(0, 0)

		if neighbours != 0 {
			t.Errorf("Expected %v but got %v", 0, neighbours)
		}
	}))

	t.Run("vertical/horizontal directions", testCase(func(t *testing.T, c *testContext) {
		game := c.game
		game.Board = [][]byte{
			{0, 1, 0},
			{1, 0, 1},
			{0, 1, 0},
		}
		neighbours := game.getNeighbours(1, 1)

		if neighbours != 4 {
			t.Errorf("Expected %v but got %v", 4, neighbours)
		}
	}))

	t.Run("diagonal directions", testCase(func(t *testing.T, c *testContext) {
		game := c.game
		game.Board = [][]byte{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		}
		neighbours := game.getNeighbours(1, 1)

		if neighbours != 4 {
			t.Errorf("Expected %v but got %v", 4, neighbours)
		}
	}))
}

func TestGame_getNextCellState(t *testing.T) {
	t.Run("lt two neighbours dies", testCase(func(t *testing.T, c *testContext) {
		game := c.game
		state := game.getNextCellState(4, 1)

		if state != byte(0) {
			t.Errorf("Expected %v but got %v", byte(0), state)
		}

		state = game.getNextCellState(4, 3)
		if state != byte(0) {
			t.Errorf("Expected %v but got %v", byte(0), state)
		}
	}))

	t.Run("two or three lives", testCase(func(t *testing.T, c *testContext) {
		game := c.game
		state := game.getNextCellState(1, 1)

		if state != byte(1) {
			t.Errorf("Expected %v but got %v", byte(1), state)
		}

		state = game.getNextCellState(1, 4)
		if state != byte(1) {
			t.Errorf("Expected %v but got %v", byte(1), state)
		}
	}))

	t.Run("two or three lives", testCase(func(t *testing.T, c *testContext) {
		game := c.game
		state := game.getNextCellState(1, 1)

		if state != byte(1) {
			t.Errorf("Expected %v but got %v", byte(1), state)
		}

		state = game.getNextCellState(1, 4)
		if state != byte(1) {
			t.Errorf("Expected %v but got %v", byte(1), state)
		}
	}))

	t.Run("gt three dies", testCase(func(t *testing.T, c *testContext) {
		game := c.game
		state := game.getNextCellState(2, 5)

		if state != byte(0) {
			t.Errorf("Expected %v but got %v", byte(0), state)
		}
	}))
}

func TestGame_Update(t *testing.T) {
	t.Run("The board is updated correctly; first figure", testCase(func(t *testing.T, c *testContext) {
		game := c.game
		game.Board = [][]byte{
			{0, 0, 1, 0, 0},
			{0, 1, 1, 0, 0},
			{0, 0, 1, 1, 0},
			{0, 0, 0, 0, 0},
		}

		game.Update(4, 5)
		expected := [][]byte{
			{0, 1, 1, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0},
		}
		if reflect.DeepEqual(game.Board, expected) {
			t.Errorf("Expected %v but got %v", expected, game.Board)
		}

		game.Update(4, 5)
		expected = [][]byte{
			{0, 1, 1, 0, 0},
			{1, 0, 0, 1, 0},
			{0, 1, 1, 0, 0},
			{0, 0, 1, 0, 0},
		}
		if reflect.DeepEqual(game.Board, expected) {
			t.Errorf("Expected %v but got %v", expected, game.Board)
		}

		game.Update(4, 5)
		expected = [][]byte{
			{0, 1, 1, 0, 0},
			{1, 0, 0, 1, 0},
			{0, 1, 1, 1, 0},
			{0, 1, 1, 0, 0},
		}
		if reflect.DeepEqual(game.Board, expected) {
			t.Errorf("Expected %v but got %v", expected, game.Board)
		}
	}))

	t.Run("The board is updated correctly; second figure (static)", testCase(func(t *testing.T, c *testContext) {
		game := c.game
		game.Board = [][]byte{
			{0, 1, 0},
			{1, 0, 1},
			{1, 0, 1},
			{0, 1, 0},
		}

		game.Update(4, 3)
		expected := [][]byte{
			{0, 1, 0},
			{1, 0, 1},
			{1, 0, 1},
			{0, 1, 0},
		}
		if reflect.DeepEqual(game.Board, expected) {
			t.Errorf("Expected %v but got %v", expected, game.Board)
		}
	}))

	t.Run("The board is updated correctly; third figure (hardcore, baby)", testCase(func(t *testing.T, c *testContext) {
		game := c.game
		game.Board = [][]byte{
			{0, 1, 0, 0},
			{1, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 0, 1, 0},
		}

		game.Update(4, 4)
		expected := [][]byte{
			{1, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 1},
			{0, 0, 0, 1},
		}
		if reflect.DeepEqual(game.Board, expected) {
			t.Errorf("Expected %v but got %v", expected, game.Board)
		}

		game.Update(4, 4)
		expected = [][]byte{
			{0, 0, 1, 0},
			{0, 0, 0, 1},
			{1, 0, 0, 0},
			{0, 1, 0, 0},
		}
		if reflect.DeepEqual(game.Board, expected) {
			t.Errorf("Expected %v but got %v", expected, game.Board)
		}

		game.Update(4, 4)
		expected = [][]byte{
			{0, 0, 0, 1},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{1, 0, 0, 0},
		}
		if reflect.DeepEqual(game.Board, expected) {
			t.Errorf("Expected %v but got %v", expected, game.Board)
		}
	}))
}
