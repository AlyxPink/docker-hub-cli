package view_explore

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/view"
	"github.com/victorbersy/docker-hub-cli/internal/ui/styles"
)

var (
	nameWidth              = 25
	organizationsnameWidth = 20
	labelsWidth            = 12
	statsWidth             = 8

	columnTitleStatsDownloads = view.ColumnTitle.Copy().Foreground(styles.DefaultTheme.StatsDownloads)
	columnTitleStatsStars     = view.ColumnTitle.Copy().Foreground(styles.DefaultTheme.StatsStars)

	ContainerPadding = 1

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)

	spinnerStyle = lipgloss.NewStyle().Padding(0, 1)

	emptyStateStyle = lipgloss.NewStyle().
			Faint(true).
			PaddingLeft(1).
			MarginBottom(1)
)
