package tabs

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/styles"
)

var (
	tabsBorderHeight  = 1
	tabsContentHeight = 2
	TabsHeight        = tabsBorderHeight + tabsContentHeight

	tab = lipgloss.NewStyle().
		Faint(true).
		Padding(0, 2)

	activeTab = tab.
			Copy().
			Faint(false).
			Bold(true).
			Background(styles.DefaultTheme.SelectedBackground).
			Foreground(styles.DefaultTheme.MainText)

	tabsRow = lipgloss.NewStyle().
		Height(tabsContentHeight).
		PaddingTop(1).
		PaddingBottom(0).
		BorderBottom(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderBottomForeground(styles.DefaultTheme.Border)

	activeView = lipgloss.NewStyle().
			Foreground(styles.DefaultTheme.MainText).
			Bold(true).
			Background(styles.DefaultTheme.SelectedBackground)

	viewsSeparator = lipgloss.NewStyle().
			BorderForeground(styles.DefaultTheme.Border).
			BorderStyle(lipgloss.NormalBorder()).
			BorderRight(true)

	inactiveView = lipgloss.NewStyle().
			Background(styles.DefaultTheme.FaintBorder).
			Foreground(styles.DefaultTheme.SecondaryText)
)
