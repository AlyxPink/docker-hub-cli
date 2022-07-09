package section

import "github.com/charmbracelet/lipgloss"

var (
	LastUpdateCellWidth = lipgloss.Width(" Updated")
	ContainerPadding    = 1

	ColumnTitle = lipgloss.NewStyle().Bold(true)

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)
)
