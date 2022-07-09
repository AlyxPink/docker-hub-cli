package listviewport

import (
	"github.com/VictorBersy/docker-hub-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	pagerHeight = 2

	pagerStyle = lipgloss.NewStyle().
			Height(pagerHeight).
			MaxHeight(pagerHeight).
			PaddingTop(1).
			Bold(true).
			Foreground(styles.DefaultTheme.FaintText)
)
