package helpers

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/Thevetat/new_bubble/constants"
)

func createFooterSpace(contentSize int) string {
	var b strings.Builder
	for i := 0; i < contentSize; i++ {
		b.WriteString("\n")
	}
	return b.String()
}

func CalculateFooterSpace(title, content string) int {
	return constants.WindowSize.Height - (lipgloss.Height(title) + lipgloss.Height(content) + 4)
}

func CreateHelpView(footerSpace int, menuItems string) string {
	space := createFooterSpace(footerSpace)
	helpMenu := constants.HelpStyle(fmt.Sprintf("\n\n ↑/↓: navigate • esc: back • q: quit %s", menuItems))
	return lipgloss.JoinVertical(lipgloss.Center, space, helpMenu)
}
