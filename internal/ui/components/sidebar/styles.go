package sidebar

import (
	"github.com/VictorBersy/docker-hub-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	borderWidth    = 1
	pagerHeight    = 2
	contentPadding = 2

	sideBarStyle = lipgloss.NewStyle().
			Padding(0, contentPadding).
			BorderLeft(true).
			BorderStyle(lipgloss.Border{
			Top:         "",
			Bottom:      "",
			Left:        "│",
			Right:       "",
			TopLeft:     "",
			TopRight:    "",
			BottomRight: "",
			BottomLeft:  "",
		}).
		BorderForeground(styles.DefaultTheme.Border)

	pagerStyle = lipgloss.NewStyle().
			Height(pagerHeight).
			PaddingTop(1).
			Bold(true).
			Foreground(styles.DefaultTheme.FaintText)
)
