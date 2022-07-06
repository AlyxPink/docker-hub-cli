package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/hack-docker-access-management-cli/internal/config"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/tabs"
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
			fmt.Println(msg)

		case key.Matches(msg, m.keys.Up):
			fmt.Println(msg)

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
	return s.String()
}

func (m *Model) getViewAt(id int) config.ViewConfig {
	views := m.ctx.Config.Views
	return views[id]
}

func (m *Model) getPrevViewId() int {
	viewsConfigs := m.ctx.Config.Views
	m.currViewId = (m.currViewId - 1) % len(viewsConfigs)
	if m.currViewId < 0 {
		m.currViewId += len(viewsConfigs)
	}

	return m.currViewId
}

func (m *Model) getNextViewId() int {
	viewsConfigs := m.ctx.Config.Views
	m.currViewId = (m.currViewId + 1) % len(viewsConfigs)
	if m.currViewId < 0 {
		m.currViewId += len(viewsConfigs)
	}

	return m.currViewId
}

func (m *Model) setCurrViewId(newViewId int) {
	m.currViewId = newViewId
	m.tabs.SetCurrViewId(newViewId)
}
