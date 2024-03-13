package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/AlyxPink/docker-hub-cli/internal/config"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/components/view"
	view_explore "github.com/AlyxPink/docker-hub-cli/internal/ui/components/view/explore"
	view_my_orgs "github.com/AlyxPink/docker-hub-cli/internal/ui/components/view/my_orgs"
	view_my_repos "github.com/AlyxPink/docker-hub-cli/internal/ui/components/view/my_repos"
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	if m.ctx.Config == nil {
		return fmt.Sprintln(m.ctx.Localizer.L("startup_reading"))
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
		mainContent = fmt.Sprintln(m.ctx.Localizer.L("startup_no_views_defined"))
	}
	s.WriteString(mainContent)
	s.WriteString("\n")
	s.WriteString(m.help.View(m.ctx))
	return s.String()
}

func (m *Model) fetchAllViews() ([]view.View, tea.Cmd) {
	explore, cmd_explore := view_explore.Fetch(0, m.ctx)
	my_repos, cmd_my_repos := view_my_repos.Fetch(1, m.ctx)
	my_orgs, cmd_my_orgs := view_my_orgs.Fetch(2, m.ctx)
	views := []view.View{explore, my_repos, my_orgs}
	cmds := []tea.Cmd{cmd_explore, cmd_my_repos, cmd_my_orgs}
	return views, tea.Batch(cmds...)
}

func (m *Model) getViews() []view.View {
	return m.views
}

func (m *Model) setViews(newViews []view.View) {
	m.views = newViews
}

func (m *Model) setCurrentView(view view.View) {
	m.currView = m.getCurrView()
	m.currViewId = view.Id()
	m.ctx.View = config.ViewType(m.currView.View())
	m.tabs.SetCurrViewId(m.currViewId)
	m.onViewedRowChanged()
}
