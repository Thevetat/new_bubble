package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Thevetat/new_bubble/tui"
)

func main() {
	if file, err := tea.LogToFile("debug.log", "debug"); err != nil {
		fmt.Println("Couldnt open file for logging: ", err)
		os.Exit(1)
	} else {
		defer func() {
			err = file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	tui.StartTea()
}
