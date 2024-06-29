package main

import (
	"fmt"
	"log"
	"os"

	"github.com/javierpoduje/2c1s/cli-client/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	c := app.NewClient(addr)
	c.Start()
}
