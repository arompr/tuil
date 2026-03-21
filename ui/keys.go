package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap = struct {
	Up           key.Binding
	Down         key.Binding
	Increase     key.Binding
	Decrease     key.Binding
	IncreaseMore key.Binding
	DecreaseMore key.Binding
}

func newKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("k/↑", "Up"),
		),
		Down: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("j/↓", "Down"),
		),
		Increase: key.NewBinding(
			key.WithKeys("l", "right"),
			key.WithHelp("l/→", "Increase"),
		),
		Decrease: key.NewBinding(
			key.WithKeys("h", "left"),
			key.WithHelp("h/←", "Decrease"),
		),
		IncreaseMore: key.NewBinding(
			key.WithKeys("L", "shift+right"),
			key.WithHelp("L/shift+→", "Increase more"),
		),
		DecreaseMore: key.NewBinding(
			key.WithKeys("H", "shift+left"),
			key.WithHelp("H/shift+←", "Decrease more"),
		),
	}
}
