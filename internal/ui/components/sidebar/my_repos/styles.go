package my_repos_sidebar

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	title     = lipgloss.NewStyle().Bold(true).MarginBottom(1)
	titleRepo = title.Copy().Margin(0, 1, 1, 0)

	statsDownloadsColors = lipgloss.AdaptiveColor{Light: "#00BBF9", Dark: "#00BBF9"}
	statsStarsColors     = lipgloss.AdaptiveColor{Light: "#FFB703", Dark: "#FFB703"}

	statsDownloads = lipgloss.NewStyle().Bold(true).PaddingRight(2).Foreground(statsDownloadsColors)
	statsStars     = lipgloss.NewStyle().Bold(true).PaddingRight(2).Foreground(statsStarsColors)
)
