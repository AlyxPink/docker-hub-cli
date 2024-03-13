package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/context"
)

type Model struct {
	CurrViewId int
}

func NewModel() Model {
	return Model{
		CurrViewId: 0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View(ctx context.ProgramContext) string {
	viewsConfigs := ctx.GetViewsConfig()
	viewTitles := make([]string, 0, len(viewsConfigs))
	for _, view := range viewsConfigs {
		viewTitles = append(viewTitles, view.Title)
	}

	var tabs []string
	for i, viewTitle := range viewTitles {
		if m.CurrViewId == i {
			tabs = append(tabs, activeTab.Render(viewTitle))
		} else {
			tabs = append(tabs, tab.Render(viewTitle))
		}
	}

	renderedTabs := lipgloss.NewStyle().
		Width(ctx.ScreenWidth).
		MaxWidth(ctx.ScreenWidth).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, tabs...))

	return tabsRow.Copy().
		Width(ctx.ScreenWidth).
		MaxWidth(ctx.ScreenWidth).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs))
}

func (m *Model) SetCurrViewId(id int) {
	m.CurrViewId = id
}
