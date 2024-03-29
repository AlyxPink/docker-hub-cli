package listviewport

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/styles"
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
