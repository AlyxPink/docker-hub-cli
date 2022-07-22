package section_explore

import (
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/section"
	"github.com/charmbracelet/lipgloss"
)

var (
	nameWidth              = 25
	organizationsnameWidth = 20
	LastUpdateCellWidth    = lipgloss.Width(" Last Update ")
	labelsWidth            = 12
	statsWidth             = 8

	statsDownloads = lipgloss.AdaptiveColor{Light: "#00BBF9", Dark: "#00BBF9"}
	statsStars     = lipgloss.AdaptiveColor{Light: "#FFB703", Dark: "#FFB703"}

	columnTitleStatsDownloads = section.ColumnTitle.Copy().Foreground(statsDownloads)
	columnTitleStatsStars     = section.ColumnTitle.Copy().Foreground(statsStars)

	ContainerPadding = 1

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)

	spinnerStyle = lipgloss.NewStyle().Padding(0, 1)

	emptyStateStyle = lipgloss.NewStyle().
			Faint(true).
			PaddingLeft(1).
			MarginBottom(1)
)