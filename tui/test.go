package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Thevetat/new_bubble/constants"
	"github.com/Thevetat/new_bubble/helpers"
)

type TestModel struct {
	quitting    bool
	footerSpace int
}

func InitTest() tea.Model {
	m := TestModel{}
	m.footerSpace = helpers.CalculateFooterSpace(constants.Title, m.contentView())
	return m
}

func (m TestModel) Init() tea.Cmd {
	return nil
}

func (m TestModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		m.footerSpace = helpers.CalculateFooterSpace(constants.Title, m.contentView())
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Main):
			return InitWelcome(), cmd
		case key.Matches(msg, constants.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit
		}
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m TestModel) headerView() string {
	return constants.Title
}

func (m TestModel) contentView() string {
	var b strings.Builder

	b.WriteString("This is some content in the test view")

	output := lipgloss.JoinHorizontal(lipgloss.Center, b.String())

	return output
}

func (m TestModel) View() string {
	if m.quitting {
		return ""
	}

	helpMenu := helpers.CreateHelpView(m.footerSpace, "• info: i • main: m")

	formatted := lipgloss.Place(
		constants.WindowSize.Width,
		constants.WindowSize.Height,
		lipgloss.Center,
		lipgloss.Top,
		lipgloss.JoinVertical(lipgloss.Center, m.headerView(), m.contentView(), helpMenu),
	)

	return constants.DocStyle.Render(formatted)
}
