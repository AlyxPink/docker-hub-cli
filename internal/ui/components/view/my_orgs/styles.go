package my_orgs

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	spinnerStyle = lipgloss.NewStyle().Padding(0, 1)

	emptyStateStyle = lipgloss.NewStyle().
			Faint(true).
			PaddingLeft(1).
			MarginBottom(1)

	ContainerPadding = 1

	containerStyle = lipgloss.NewStyle().
			Padding(0, ContainerPadding)
)
