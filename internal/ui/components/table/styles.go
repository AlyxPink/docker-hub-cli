package table

import (
	"github.com/VictorBersy/docker-hub-cli/internal/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	cellStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			MaxHeight(1)

	selectedCellStyle = cellStyle.Copy().
				Background(styles.DefaultTheme.SelectedBackground)

	titleCellStyle = cellStyle.Copy().
			Bold(true).
			Foreground(styles.DefaultTheme.MainText)

	singleRuneTitleCellStyle = titleCellStyle.Copy().Width(styles.SingleRuneWidth)

	headerStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(styles.DefaultTheme.SecondaryBorder).
			BorderBottom(true)

	rowStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(styles.DefaultTheme.FaintBorder).
			BorderBottom(true).
			MaxHeight(1)
)
