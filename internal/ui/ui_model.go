package ui

import (
	"fmt"
	"log"

	"github.com/VictorBersy/docker-hub-cli/internal/config"
	"github.com/VictorBersy/docker-hub-cli/internal/data"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/help"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/sidebar"
	sidebar_repository "github.com/VictorBersy/docker-hub-cli/internal/ui/components/sidebar/repository"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/tabs"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/view"
	"github.com/VictorBersy/docker-hub-cli/internal/ui/context"
	"github.com/VictorBersy/docker-hub-cli/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
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
			Config: &config.Config{},
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
			if prevView != nil {
				log.Println("prevView")
				m.ctx.View = m.switchSelectedView()
				m.setCurrViewId(prevView.Id())
				m.syncMainContentWidth() // TODO check if it's required
				m.setCurrentView(prevView)
				m.onViewedRowChanged()
			}

		case key.Matches(msg, m.keys.NextView):
			nextViewId := m.getNextViewId()
			nextView := m.getViewAt(nextViewId)
			if nextView != nil {
				log.Println("nextView")
				m.ctx.View = m.switchSelectedView()
				m.setCurrViewId(nextView.Id())
				m.syncMainContentWidth() // TODO check if it's required
				m.setCurrentView(nextView)
				m.onViewedRowChanged()
			}

		case key.Matches(msg, m.keys.NextView):
			fmt.Print()

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
			log.Println()
			cmd = tea.Quit

		}

	case initMsg:
		m.ctx.Config = &msg.Config
		m.ctx.View = m.ctx.Config.Defaults.View
		m.sidebar.IsOpen = msg.Config.Defaults.Preview.Open
		m.syncMainContentWidth()
		newViews, fetchViewsCmds := m.fetchAllViews()
		m.setViews(newViews)
		cmd = fetchViewsCmds

	case view.ViewMsg:
		cmd = m.updateRelevantView(msg)

		if msg.GetViewId() == m.currViewId {
			m.onViewedRowChanged()
		}

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

func (m *Model) setCurrViewId(newViewId int) {
	m.currViewId = newViewId
	m.tabs.SetCurrViewId(newViewId)
}

func (m *Model) onViewedRowChanged() {
	m.syncSidebarExplore()
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

func (m *Model) syncMainContentWidth() {
	sideBarOffset := 0
	if m.sidebar.IsOpen {
		sideBarOffset = m.ctx.Config.Defaults.Preview.Width
	}
	m.ctx.MainContentWidth = m.ctx.ScreenWidth - sideBarOffset
}

func (m *Model) syncSidebarExplore() {
	currRowData := m.getCurrRowData()
	width := m.sidebar.GetSidebarContentWidth()

	switch row_data := currRowData.(type) {
	case *data.RepositoryData:
		content := sidebar_repository.NewModel(row_data, width).View()
		m.sidebar.SetContent(content)
	}
}
