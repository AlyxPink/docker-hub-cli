package help

import (
	"github.com/VictorBersy/docker-hub-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	FooterHeight = 3

	helpTextStyle = lipgloss.NewStyle().Foreground(styles.DefaultTheme.SecondaryText)
	helpStyle     = lipgloss.NewStyle().
			Height(FooterHeight - 1).
			BorderTop(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(styles.DefaultTheme.Border)
)
