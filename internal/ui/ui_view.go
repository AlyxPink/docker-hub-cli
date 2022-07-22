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
	if m.ctx.View == config.ExploreView {
		return view_explore.FetchAllViews(m.ctx)
	} else {
		return view_my_repos.FetchAllViews(m.ctx)
	}
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
}

func (m *Model) switchSelectedView() config.ViewType {
	// TODO: add multiple views
	if m.ctx.View == config.ExploreView {
		return config.MyReposView
	} else {
		return config.ExploreView
	}
}
