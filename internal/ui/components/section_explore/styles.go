package section_explore

import (
	"github.com/VictorBersy/docker-hub-cli/internal/ui/components/section"
	"github.com/charmbracelet/lipgloss"
)

var (
	nameWidth              = 25
	organizationsnameWidth = 20
	LastUpdateCellWidth    = lipgloss.Width(" Last Update ")
	labelWidth             = 4
	statsWidth             = 8

	labelDockerOfficial    = lipgloss.AdaptiveColor{Light: "#2E7F74", Dark: "#2E7F74"}
	labelVerifiedPublisher = lipgloss.AdaptiveColor{Light: "#086DD7", Dark: "#086DD7"}
	labelOpenSourceProgram = lipgloss.AdaptiveColor{Light: "#7D2EFF", Dark: "#7D2EFF"}
	statsDownloads         = lipgloss.AdaptiveColor{Light: "#00BBF9", Dark: "#00BBF9"}
	statsStars             = lipgloss.AdaptiveColor{Light: "#FFB703", Dark: "#FFB703"}

	columnTitleLabelDockerOfficial    = section.ColumnTitle.Copy().Foreground(labelDockerOfficial)
	columnTitleLabelVerifiedPublisher = section.ColumnTitle.Copy().Foreground(labelVerifiedPublisher)
	columnTitleLabelOpenSourceProgram = section.ColumnTitle.Copy().Foreground(labelOpenSourceProgram)
	columnTitleStatsDownloads         = section.ColumnTitle.Copy().Foreground(statsDownloads)
	columnTitleStatsStars             = section.ColumnTitle.Copy().Foreground(statsStars)

	ContainerPadding = 1

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)

	spinnerStyle = lipgloss.NewStyle().Padding(0, 1)

	emptyStateStyle = lipgloss.NewStyle().
			Faint(true).
			PaddingLeft(1).
			MarginBottom(1)
)
