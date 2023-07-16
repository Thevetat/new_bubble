package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Thevetat/new_bubble/constants"
)

func StartTea() error {
	model := InitWelcome()
	constants.Program = tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if err := constants.Program.Start(); err != nil {
		fmt.Println("Couldnt start program: ", err)
		os.Exit(1)
	}

	return nil
}
