package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Thevetat/new_bubble/constants"
	"github.com/Thevetat/new_bubble/helpers"
)

type Status struct{}

type WelcomeModel struct {
	status      string
	isFkd       bool
	quitting    bool
	footerSpace int
}

func InitWelcome() tea.Model {
	m := WelcomeModel{}
	m.footerSpace = helpers.CalculateFooterSpace(m.headerView(), m.contentView())
	return m
}

func (m WelcomeModel) Init() tea.Cmd {
	return nil
}

func (m WelcomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		m.footerSpace = helpers.CalculateFooterSpace(m.headerView(), m.contentView())
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Test):
			return InitTest(), cmd
		case key.Matches(msg, constants.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit
		}
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m WelcomeModel) headerView() string {
	return constants.Title
}

func (m WelcomeModel) contentView() string {
	var b strings.Builder

	b.WriteString("This is some content")

	output := lipgloss.JoinHorizontal(lipgloss.Center, b.String())

	return output
}

func (m WelcomeModel) View() string {
	if m.quitting {
		return ""
	}

	helpMenu := helpers.CreateHelpView(m.footerSpace, "• info: i • status: s")

	formatted := lipgloss.Place(
		constants.WindowSize.Width,
		constants.WindowSize.Height,
		lipgloss.Center,
		lipgloss.Top,
		lipgloss.JoinVertical(lipgloss.Center, m.headerView(), m.contentView(), helpMenu),
	)

	return constants.DocStyle.Render(formatted)
}
