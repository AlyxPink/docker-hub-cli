package section

import "github.com/charmbracelet/lipgloss"

var (
	LastUpdateCellWidth = lipgloss.Width(" Updated")
	ContainerPadding    = 1

	ColumnTitle = lipgloss.NewStyle().Bold(true)
	LabelTitle  = lipgloss.NewStyle().Width(3)

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)
)
