package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/context"
)

type Model struct {
	currSectionId int
}

func NewModel() Model {
	return Model{
		currSectionId: 0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) Section(ctx context.ProgramContext) string {
	sections := ctx.Config.Sections
	var tabs []string
	for i, section := range sections {
		title := section.Title()
		if m.currSectionId == i {
			tabs = append(tabs, activeTab.Render(title))
		} else {
			tabs = append(tabs, tab.Render(title))
		}
	}

	tabsWidth := ctx.ScreenWidth
	renderedTabs := lipgloss.NewStyle().
		Width(tabsWidth).
		MaxWidth(tabsWidth).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, tabs...))

	return tabsRow.Copy().
		Width(ctx.ScreenWidth).
		MaxWidth(ctx.ScreenWidth).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs))
}

func (m *Model) SetCurrSectionId(id int) {
	m.currSectionId = id
}
