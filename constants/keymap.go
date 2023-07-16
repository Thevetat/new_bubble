package constants

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	Main         key.Binding
	Info         key.Binding
	Status       key.Binding
	Test         key.Binding
	Enter        key.Binding
	CreateConfig key.Binding
	Rename       key.Binding
	Delete       key.Binding
	Back         key.Binding
	Quit         key.Binding
	Up           key.Binding
	Down         key.Binding
}

// Keymap reusable key mappings shared across models
var Keymap = keymap{
	Main: key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "main"),
	),
	Info: key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", "info"),
	),
	Status: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "status"),
	),
	Test: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "test"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Rename: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "rename"),
	),
	CreateConfig: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "create config"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	),
}
