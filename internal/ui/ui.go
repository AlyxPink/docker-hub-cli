package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/hack-docker-access-management-cli/internal/config"
	"github.com/docker/hack-docker-access-management-cli/internal/data"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/help"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/section"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/section_explore"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/sidebar"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/sidebar_repository"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/tabs"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/context"
	"github.com/docker/hack-docker-access-management-cli/internal/utils"
)

type Model struct {
	keys          utils.KeyMap
	err           error
	sidebar       sidebar.Model
	currSectionId int
	help          help.Model
	explore       []section.Section
	tabs          tabs.Model
	ctx           context.ProgramContext
}

func NewModel() Model {
	tabsModel := tabs.NewModel()
	return Model{
		keys:          utils.Keys,
		help:          help.NewModel(),
		currSectionId: 0,
		tabs:          tabsModel,
		sidebar:       sidebar.NewModel(),
	}
}

func initScreen() tea.Msg {
	return initMsg{Config: config.GetDefaultConfig()}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(initScreen, tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		sidebarCmd  tea.Cmd
		helpCmd     tea.Cmd
		cmds        []tea.Cmd
		currSection = m.getCurrSection()
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.PrevSection):
			prevSection := m.getSectionAt(m.getPrevSectionId())
			if prevSection != nil {
				m.setCurrSectionId(prevSection.Id())
				m.onViewedRowChanged()
			}

		case key.Matches(msg, m.keys.NextSection):
			nextSectionId := m.getNextSectionId()
			nextSection := m.getSectionAt(nextSectionId)
			if nextSection != nil {
				m.setCurrSectionId(nextSection.Id())
				m.onViewedRowChanged()
			}

		case key.Matches(msg, m.keys.Down):
			currSection.NextRow()
			m.onViewedRowChanged()

		case key.Matches(msg, m.keys.Up):
			currSection.PrevRow()
			m.onViewedRowChanged()

		case key.Matches(msg, m.keys.FirstLine):
			currSection.FirstItem()
			m.onViewedRowChanged()

		case key.Matches(msg, m.keys.LastLine):
			currSection.LastItem()
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
			cmd = currSection.FetchSectionRows()

		case key.Matches(msg, m.keys.Quit):
			cmd = tea.Quit

		}

	case initMsg:
		m.ctx.Config = &msg.Config
		m.ctx.View = m.ctx.Config.Defaults.View
		m.sidebar.IsOpen = msg.Config.Defaults.Preview.Open
		m.syncMainContentWidth()
		newSections, fetchSectionsCmds := m.fetchAllViewSections()
		m.setCurrentViewSections(newSections)
		cmd = fetchSectionsCmds

	case section.SectionMsg:
		cmd = m.updateRelevantSection(msg)

		if msg.GetSectionId() == m.currSectionId {
			switch msg.GetSectionType() {
			case section_explore.SectionType:
				m.onViewedRowChanged()
			}
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
	currSection := m.getCurrSection()
	mainContent := ""
	if currSection != nil {
		mainContent = lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.getCurrSection().View(),
			m.sidebar.View(),
		)
	} else {
		mainContent = "No sections defined..."
	}
	s.WriteString(mainContent)
	s.WriteString("\n")
	s.WriteString(m.help.View(m.ctx))
	return s.String()
}

type initMsg struct {
	Config config.Config
}

type errMsg struct {
	error
}

func (e errMsg) Error() string { return e.error.Error() }

func (m *Model) setCurrSectionId(newSectionId int) {
	m.currSectionId = newSectionId
	m.tabs.SetCurrSectionId(newSectionId)
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
	for _, section := range m.getCurrentViewSections() {
		section.UpdateProgramContext(&m.ctx)
	}
	m.sidebar.UpdateProgramContext(&m.ctx)
}

func (m *Model) updateRelevantSection(msg section.SectionMsg) (cmd tea.Cmd) {
	var updatedSection section.Section

	switch msg.GetSectionType() {
	case section_explore.SectionType:
		updatedSection, cmd = m.explore[msg.GetSectionId()].Update(msg)
		m.explore[msg.GetSectionId()] = updatedSection
	}

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

func (m *Model) fetchAllViewSections() ([]section.Section, tea.Cmd) {
	return section_explore.FetchAllSections(m.ctx)
}

func (m *Model) getCurrentViewSections() []section.Section {
	return m.explore
}

func (m *Model) setCurrentViewSections(newSections []section.Section) {
	if m.ctx.View == config.ExploreView {
		m.explore = newSections
	}
}

func (m *Model) switchSelectedView() config.ViewType {
	return config.ExploreView
}
