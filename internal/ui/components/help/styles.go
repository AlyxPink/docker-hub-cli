package help

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/hack-docker-access-management-cli/internal/ui/styles"
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
