package my_repos_sidebar

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	visibilityPrivate = lipgloss.NewStyle().Foreground(lipgloss.Color("#EF476F")).Padding(1, 0, 2, 0)
	visibilityPublic  = lipgloss.NewStyle().Foreground(lipgloss.Color("#06D6A0")).Padding(1, 0, 2, 0)

	statsDownloadsColors = lipgloss.AdaptiveColor{Light: "#00BBF9", Dark: "#00BBF9"}
	statsStarsColors     = lipgloss.AdaptiveColor{Light: "#FFB703", Dark: "#FFB703"}

	statsDownloads = lipgloss.NewStyle().Bold(true).PaddingRight(2).Foreground(statsDownloadsColors)
	statsStars     = lipgloss.NewStyle().Bold(true).PaddingRight(2).Foreground(statsStarsColors)
)
