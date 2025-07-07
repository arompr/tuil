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
			key.WithKeys("k"),
			key.WithHelp("k", "Up"),
		),
		Down: key.NewBinding(
			key.WithKeys("j"),
			key.WithHelp("j", "Down"),
		),
		Increase: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "Increase"),
		),
		Decrease: key.NewBinding(
			key.WithKeys("h"),
			key.WithHelp("h", "Decrease"),
		),
		IncreaseMore: key.NewBinding(
			key.WithKeys("<S-l>"),
			key.WithHelp("<S-l>", "Increase more"),
		),
		DecreaseMore: key.NewBinding(
			key.WithKeys("<S-h>"),
			key.WithHelp("<S-h>", "Decrease more"),
		),
	}
}
