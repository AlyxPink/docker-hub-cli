package my_repos

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/view"
	"github.com/victorbersy/docker-hub-cli/internal/ui/styles"
)

var (
	isPrivateWidth = 4
	updatedAtWidth = lipgloss.Width(" Last Update ")
	createdAtWidth = lipgloss.Width(" Created at ")
	statsWidth     = 8

	columnTitleIsPrivate      = view.ColumnTitle.Copy().Foreground(styles.DefaultTheme.MainText)
	columnTitleStatsDownloads = view.ColumnTitle.Copy().Foreground(styles.DefaultTheme.StatsDownloads)
	columnTitleStatsStars     = view.ColumnTitle.Copy().Foreground(styles.DefaultTheme.StatsStars)

	spinnerStyle = lipgloss.NewStyle().Padding(0, 1)

	emptyStateStyle = lipgloss.NewStyle().
			Faint(true).
			PaddingLeft(1).
			MarginBottom(1)

	ContainerPadding = 1

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)
)
