package constants

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	Program    *tea.Program
	WindowSize tea.WindowSizeMsg
)

var (
	DocStyle       = lipgloss.NewStyle().Margin(0, 2)
	PrimaryStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#1abc9c")).Render
	SecondaryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bb9af7")).Render
	TertiaryStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#399a96")).Render
	HelpStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#73daca")).Render
	InfoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#7dcfff")).Render
	SuccStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#9ece6a")).Render
	WarnStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff9e64")).Render
	ErrStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#f7768e")).Render
	BlurredStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#f0f0f0"))
	QuitTextStyle  = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

var (
	InputFocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#d074ec"))
	InputCursorStyle  = InputFocusedStyle.Copy()
	InputNoStyle      = lipgloss.NewStyle()
	FocusedButton     = InputFocusedStyle.Copy().Render("[ Submit ]")
	BlurredButton     = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Submit"))
)

var (
	ListTitleStyle        = lipgloss.NewStyle().MarginLeft(2)
	ListItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	ListSelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	ListPaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	ListHelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

var (
	InfoRStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return lipgloss.NewStyle().BorderStyle(b)
	}()

	InfoLStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b)
	}()
)

func StyleBool(s string, b bool, e bool) string {
	enabled := "False"
	if e {
		enabled = "True"
	}
	if b {
		return SuccStyle(fmt.Sprintf("%s - Enabled: %v", s, enabled))
	} else {
		return ErrStyle(fmt.Sprintf("%s - Enabled: %v", s, enabled))
	}
}

func CenterView(val string) string {
	return lipgloss.Place(
		WindowSize.Width,
		WindowSize.Height,
		lipgloss.Center,
		lipgloss.Center,
		val,
	)
}
