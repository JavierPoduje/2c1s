package main

import (
	"log"

	"github.com/javierpoduje/2c1s/server/server"
	"github.com/joho/godotenv"
)

const (
	BoardWidth  = 4
	BoardHeight = 4
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := server.NewServer(BoardWidth, BoardHeight)
	s.Start()
}
