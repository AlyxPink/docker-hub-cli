package repository

import "github.com/charmbracelet/lipgloss"

var (
	EmptyLabel     = lipgloss.NewStyle().Width(3).Render("")
	LabelWithGlyph = lipgloss.NewStyle().Width(3)

	VisibilityPrivate = lipgloss.NewStyle().Foreground(lipgloss.Color("#EF476F"))
	VisibilityPublic  = lipgloss.NewStyle().Foreground(lipgloss.Color("#06D6A0"))
)
