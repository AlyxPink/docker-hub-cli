package ui

import (
	"strings"

	"github.com/VictorBersy/docker-hub-cli/internal/config"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/view"
	view_explore "github.com/VictorBersy/docker-hub-cli/internal/ui/components/view/explore"
	view_my_repos "github.com/VictorBersy/docker-hub-cli/internal/ui/components/view/my_repos"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	if m.ctx.Config == nil {
		return "Reading config...\n"
	}

	s := strings.Builder{}
	s.WriteString(m.tabs.View(m.ctx))
	s.WriteString("\n")
	mainContent := ""

	if m.currView != nil {
		mainContent = lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.getCurrView().View(),
			m.sidebar.View(),
		)
	} else {
		mainContent = "No views defined..."
	}
	s.WriteString(mainContent)
	s.WriteString("\n")
	s.WriteString(m.help.View(m.ctx))
	return s.String()
}

func (m *Model) fetchAllViews() ([]view.View, tea.Cmd) {
	explore, cmd_explore := view_explore.Fetch(m.ctx)
	my_repos, cmd_my_repos := view_my_repos.Fetch(m.ctx)
	views := []view.View{explore, my_repos}
	cmds := []tea.Cmd{cmd_explore, cmd_my_repos}
	return views, tea.Batch(cmds...)
}

func (m *Model) getViews() []view.View {
	return m.views
}

func (m *Model) setViews(newViews []view.View) {
	// TODO: add multiple views
	if m.ctx.View == config.ExploreView {
		m.views = newViews
	} else {
		m.views = newViews
	}
}

func (m *Model) setCurrentView(view view.View) {
	m.currView = m.getCurrView()
	m.currViewId = view.Id()
	m.tabs.SetCurrViewId(m.currViewId)
	m.onViewedRowChanged()
}

func (m *Model) switchSelectedView() config.ViewType {
	// TODO: add multiple views
	if m.ctx.View == config.ExploreView {
		return config.MyReposView
	} else {
		return config.ExploreView
	}
}
