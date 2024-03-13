package my_repos_sidebar

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/styles"
)

var (
	visibilityPrivate = lipgloss.NewStyle().Foreground(styles.DefaultTheme.ErrorText).Padding(1, 0, 2, 0)
	visibilityPublic  = lipgloss.NewStyle().Foreground(styles.DefaultTheme.SuccessText).Padding(1, 0, 2, 0)

	statsDownloads = lipgloss.NewStyle().Bold(true).PaddingRight(2).Foreground(styles.DefaultTheme.StatsDownloads)
	statsStars     = lipgloss.NewStyle().Bold(true).PaddingRight(2).Foreground(styles.DefaultTheme.StatsStars)
)
