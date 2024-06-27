package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	StartClient()
}

// StartClient starts the client and listens for messages from the server
func StartClient() {
	conn, err := net.Dial("tcp", os.Getenv("ADDRESS"))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server at", os.Getenv("ADDRESS"))

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		printMessage(buffer[:n])
	}
}

func printMessage(message []byte) {
	width := int(message[0])
	height := int(message[1])
	board := message[2:]

	str := strings.Builder{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if board[y*width+x] == 0 {
				str.WriteString("-")
			} else {
				str.WriteString("X")
			}
		}
		str.WriteString("\n")
	}

	fmt.Println("mirame baby...")
	fmt.Println(str.String())
}
