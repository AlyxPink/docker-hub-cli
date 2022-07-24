package my_repos

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/view"
)

var (
	isPrivateWidth = 4
	updatedAtWidth = lipgloss.Width(" Last Update ")
	createdAtWidth = lipgloss.Width(" Created at ")
	statsWidth     = 8

	isPrivate      = lipgloss.AdaptiveColor{Light: "#edede9", Dark: "#edede9"}
	statsDownloads = lipgloss.AdaptiveColor{Light: "#00BBF9", Dark: "#00BBF9"}
	statsStars     = lipgloss.AdaptiveColor{Light: "#FFB703", Dark: "#FFB703"}

	columnTitleIsPrivate      = view.ColumnTitle.Copy().Foreground(isPrivate)
	columnTitleStatsDownloads = view.ColumnTitle.Copy().Foreground(statsDownloads)
	columnTitleStatsStars     = view.ColumnTitle.Copy().Foreground(statsStars)

	spinnerStyle = lipgloss.NewStyle().Padding(0, 1)

	emptyStateStyle = lipgloss.NewStyle().
			Faint(true).
			PaddingLeft(1).
			MarginBottom(1)

	ContainerPadding = 1

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)
)
