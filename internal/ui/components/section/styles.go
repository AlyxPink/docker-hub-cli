package section

import "github.com/charmbracelet/lipgloss"

var (
	LastUpdateCellWidth = lipgloss.Width(" Updated")
	ContainerPadding    = 1

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)
)
