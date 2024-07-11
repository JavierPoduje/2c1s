package main

import (
	"log"

	"github.com/javierpoduje/2c1s/server/sender"
	"github.com/joho/godotenv"
)

const (
	BoardWidth  = 13
	BoardHeight = 13
)

// place the diamong right at the top-left of the board
func diamondRing(height, width int) [][]int {
	if width < 13 || height < 13 {
		log.Panic("Board must be at least 13x13 to display the diamond-ring")
	}

	return [][]int{
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1},
		{0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	}
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	seed := diamondRing(BoardHeight, BoardWidth)

	s := sender.NewServer(BoardWidth, BoardHeight, seed)
	s.Start(BoardWidth, BoardHeight, seed)
}
