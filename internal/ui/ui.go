package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/hack-docker-access-management-cli/internal/config"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/tabs"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/view"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/context"
	"github.com/docker/hack-docker-access-management-cli/internal/utils"
)

type Model struct {
	keys       utils.KeyMap
	tabs       tabs.Model
	currViewId int
	ctx        context.ProgramContext
}

func NewModel() Model {
	return Model{
		keys: utils.Keys,
		tabs: tabs.NewModel(),
		ctx: context.ProgramContext{
			Config: config.DefaultConfig(),
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.PrevView):
			view := m.getViewAt(m.getPrevViewId())
			m.setCurrViewId(view.View.Id)

		case key.Matches(msg, m.keys.NextView):
			view := m.getViewAt(m.getNextViewId())
			m.setCurrViewId(view.View.Id)

		case key.Matches(msg, m.keys.Down):

		case key.Matches(msg, m.keys.Up):

		case key.Matches(msg, m.keys.Quit):
			cmd = tea.Quit

		}
	case tea.WindowSizeMsg:
		m.ctx.ScreenWidth, m.ctx.ScreenHeight = msg.Width, msg.Height
		return m, nil
	}
	return m, tea.Batch(cmd)
}

func (m Model) View() string {
	s := strings.Builder{}
	s.WriteString(m.tabs.View(m.ctx))
	s.WriteString("\n")
	mainContent := ""
	mainContent = lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.getCurrView().View(),
	)
	s.WriteString(mainContent)
	s.WriteString("\n")
	return s.String()
}

func (m *Model) getViewAt(id int) config.ViewConfig {
	views := m.ctx.Config.Views
	return views[id]
}

func (m *Model) getPrevViewId() int {
	views := m.ctx.Config.Views
	m.currViewId = (m.currViewId - 1) % len(views)
	if m.currViewId < 0 {
		m.currViewId += len(views)
	}

	return m.currViewId
}

func (m *Model) getNextViewId() int {
	views := m.ctx.Config.Views
	m.currViewId = (m.currViewId + 1) % len(views)
	if m.currViewId < 0 {
		m.currViewId += len(views)
	}

	return m.currViewId
}

func (m *Model) setCurrViewId(newViewId int) {
	m.currViewId = newViewId
	m.tabs.SetCurrViewId(newViewId)
}

func (m *Model) getCurrView() view.Model {
	views := m.ctx.Config.Views
	return views[m.currViewId].View
}
