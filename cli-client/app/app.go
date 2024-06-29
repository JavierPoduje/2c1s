package app

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (c *Client) Start() {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server at", c.addr)

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
				str.WriteString("- ")
			} else {
				str.WriteString("X ")
			}
		}
		str.WriteString("\n")
	}

	fmt.Printf("(%v, %v)\n", width, height)
	fmt.Println(str.String())
}
