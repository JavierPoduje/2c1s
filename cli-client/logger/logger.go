package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Logger struct {
	filename string
}

func NewLogger(filename string) *Logger {
	return &Logger{
		filename: filename,
	}
}

func (l Logger) Log(msg string) {
	f, err := tea.LogToFile(l.filename, "")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	log.SetFlags(log.Lmsgprefix)
	now := time.Now()
	formattedTime := now.Format("15:04:05")

	log.Printf("%s", fmt.Sprintf("[%s]: %s", formattedTime, msg))
}
