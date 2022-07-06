package utils

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up        key.Binding
	Down      key.Binding
	FirstLine key.Binding
	LastLine  key.Binding
	PageDown  key.Binding
	PageUp    key.Binding
	NextView  key.Binding
	PrevView  key.Binding
	Help      key.Binding
	Quit      key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.FirstLine, k.LastLine},
		{k.PrevView, k.NextView},
		{k.PageDown, k.PageUp},
		{k.Help, k.Quit},
	}
}

var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	PrevView: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("/h", "previous section"),
	),
	NextView: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("/l", "next section"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
