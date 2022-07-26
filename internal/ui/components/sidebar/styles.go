package sidebar

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/victorbersy/docker-hub-cli/internal/ui/styles"
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

	Title         = lipgloss.NewStyle().Bold(true).MarginBottom(1)
	SubTitle      = Title.Copy().Margin(0, 1, 1, 0)
	AttributeName = lipgloss.NewStyle().Bold(true)
	TextBox       = lipgloss.NewStyle().Margin(1, 0, 2, 0)
)
