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

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := sender.NewServer(BoardWidth, BoardHeight)
	s.Start(BoardWidth, BoardHeight)
}
