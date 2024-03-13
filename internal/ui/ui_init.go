package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/AlyxPink/docker-hub-cli/internal/config"
)

func initScreen() tea.Msg {
	return initMsg{Config: config.GetDefaultConfig()}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(initScreen, tea.EnterAltScreen)
}

type initMsg struct {
	Config config.Config
}

type errMsg struct {
	error
}

func (e errMsg) Error() string { return e.error.Error() }
