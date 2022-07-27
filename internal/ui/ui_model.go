package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/victorbersy/docker-hub-cli/internal/config"
	"github.com/victorbersy/docker-hub-cli/internal/config/locales"
	data_search "github.com/victorbersy/docker-hub-cli/internal/data/search"
	data_user "github.com/victorbersy/docker-hub-cli/internal/data/user"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/help"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/sidebar"
	explore_sidebar "github.com/victorbersy/docker-hub-cli/internal/ui/components/sidebar/explore"
	my_repos_sidebar "github.com/victorbersy/docker-hub-cli/internal/ui/components/sidebar/my_repos"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/tabs"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/view"
	"github.com/victorbersy/docker-hub-cli/internal/ui/context"
	"github.com/victorbersy/docker-hub-cli/internal/utils"
)

type Model struct {
	keys       utils.KeyMap
	err        error
	sidebar    sidebar.Model
	currViewId int
	currView   view.View
	help       help.Model
	views      []view.View
	tabs       tabs.Model
	ctx        context.ProgramContext
}

func NewModel() Model {
	tabsModel := tabs.NewModel()
	return Model{
		keys: utils.Keys,
		help: help.NewModel(),
		tabs: tabsModel,
		ctx: context.ProgramContext{
			Config:    &config.Config{},
			Localizer: locales.GetLocalizer(),
		},
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd        tea.Cmd
		sidebarCmd tea.Cmd
		helpCmd    tea.Cmd
		cmds       []tea.Cmd
		currView   = m.getCurrView()
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.PrevView):
			prevView := m.getViewAt(m.getPrevViewId())
			m.ctx.View = m.switchSelectedView()
			m.setCurrentView(prevView)

		case key.Matches(msg, m.keys.NextView):
			nextView := m.getViewAt(m.getNextViewId())
			m.ctx.View = m.switchSelectedView()
			m.setCurrentView(nextView)

		case key.Matches(msg, m.keys.Down):
			currView.NextRow()
			m.onViewedRowChanged()

		case key.Matches(msg, m.keys.Up):
			currView.PrevRow()
			m.onViewedRowChanged()

		case key.Matches(msg, m.keys.FirstLine):
			currView.FirstItem()
			m.onViewedRowChanged()

		case key.Matches(msg, m.keys.LastLine):
			currView.LastItem()
			m.onViewedRowChanged()

		case key.Matches(msg, m.keys.TogglePreview):
			m.sidebar.IsOpen = !m.sidebar.IsOpen
			m.syncMainContentWidth()

		case key.Matches(msg, m.keys.OpenDockerHub):
			currRow := m.getCurrRowData()
			if currRow != nil {
				utils.OpenBrowser(currRow.GetUrl())
			}

		case key.Matches(msg, m.keys.Refresh):
			cmd = currView.FetchViewRows()

		case key.Matches(msg, m.keys.Quit):
			cmd = tea.Quit

		}

	case initMsg:
		m.ctx.Config = &msg.Config
		m.ctx.View = m.ctx.Config.Defaults.View
		m.sidebar.IsOpen = msg.Config.Defaults.Preview.Open
		m.syncMainContentWidth()
		newViews, fetchViewsCmds := m.fetchAllViews()
		m.setViews(newViews)
		m.setCurrentView(m.views[0])
		cmd = fetchViewsCmds

	case view.ViewMsg:
		cmd = m.updateRelevantView(msg)
		m.onViewedRowChanged()

	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)

	case errMsg:
		m.err = msg
	}

	m.syncProgramContext()
	m.sidebar, sidebarCmd = m.sidebar.Update(msg)
	m.help, helpCmd = m.help.Update(msg)
	cmds = append(cmds, cmd, sidebarCmd, helpCmd)
	return m, tea.Batch(cmds...)
}

func (m *Model) onViewedRowChanged() {
	m.syncSidebar()
}

func (m *Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.help.SetWidth(msg.Width)
	m.ctx.ScreenWidth = msg.Width
	m.ctx.ScreenHeight = msg.Height
	m.ctx.MainContentHeight = msg.Height - tabs.TabsHeight - help.FooterHeight
	m.syncMainContentWidth()
}

func (m *Model) syncProgramContext() {
	for _, view := range m.getViews() {
		view.UpdateProgramContext(&m.ctx)
	}
	m.sidebar.UpdateProgramContext(&m.ctx)
}

func (m *Model) updateRelevantView(msg view.ViewMsg) (cmd tea.Cmd) {
	var updatedView view.View

	updatedView, cmd = m.views[msg.GetViewId()].Update(msg)
	m.views[msg.GetViewId()] = updatedView

	return cmd
}

func (m *Model) syncSidebar() {
	currRowData := m.getCurrRowData()
	width := m.sidebar.GetSidebarContentWidth()
	switch data := currRowData.(type) {
	case *data_search.Repository:
		content := explore_sidebar.NewModel(data, width).View()
		m.sidebar.SetContent(content)
	case *data_user.Repository:
		content := my_repos_sidebar.NewModel(data, width).View()
		m.sidebar.SetContent(content)
	}
}

func (m *Model) syncMainContentWidth() {
	sideBarOffset := 0
	if m.sidebar.IsOpen {
		sideBarOffset = m.ctx.Config.Defaults.Preview.Width
	}
	m.ctx.MainContentWidth = m.ctx.ScreenWidth - sideBarOffset
}
