package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/hack-docker-access-management-cli/internal/config"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/section"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/tabs"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/context"
	"github.com/docker/hack-docker-access-management-cli/internal/utils"
)

type Model struct {
	keys          utils.KeyMap
	tabs          tabs.Model
	currSectionId int
	ctx           context.ProgramContext
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
		case key.Matches(msg, m.keys.PrevSection):
			Section := m.getSectionAt(m.getPrevSectionId())
			m.setCurrSectionId(Section.Section.Id)

		case key.Matches(msg, m.keys.NextSection):
			Section := m.getSectionAt(m.getNextSectionId())
			m.setCurrSectionId(Section.Section.Id)

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
	s.WriteString(m.tabs.Section(m.ctx))
	s.WriteString("\n")
	mainContent := ""
	mainContent = lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.getCurrSection().View(),
	)
	s.WriteString(mainContent)
	s.WriteString("\n")
	return s.String()
}

func (m *Model) getSectionAt(id int) config.SectionConfig {
	Sections := m.ctx.Config.Sections
	return Sections[id]
}

func (m *Model) getPrevSectionId() int {
	Sections := m.ctx.Config.Sections
	m.currSectionId = (m.currSectionId - 1) % len(Sections)
	if m.currSectionId < 0 {
		m.currSectionId += len(Sections)
	}

	return m.currSectionId
}

func (m *Model) getNextSectionId() int {
	Sections := m.ctx.Config.Sections
	m.currSectionId = (m.currSectionId + 1) % len(Sections)
	if m.currSectionId < 0 {
		m.currSectionId += len(Sections)
	}

	return m.currSectionId
}

func (m *Model) setCurrSectionId(newSectionId int) {
	m.currSectionId = newSectionId
	m.tabs.SetCurrSectionId(newSectionId)
}

func (m *Model) getCurrSection() section.Model {
	Sections := m.ctx.Config.Sections
	return Sections[m.currSectionId].Section
}
