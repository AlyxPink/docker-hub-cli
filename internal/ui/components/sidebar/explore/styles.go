package explore_sidebar

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	archTagFg          = lipgloss.AdaptiveColor{Light: "#22223b", Dark: "#EDEDE9"}
	archTagBg          = lipgloss.AdaptiveColor{Light: "#bde0fe", Dark: "#0077b6"}
	dockerPullCmdBoxBg = lipgloss.AdaptiveColor{Light: "#00f5d4", Dark: "#8ac926"}
	dockerPullCmdBoxFg = lipgloss.Color("#22223b")

	title      = lipgloss.NewStyle().Bold(true).MarginBottom(1)
	titleRepo  = title.Copy().Margin(0, 1, 1, 0)
	labelGlyph = title.Copy().Margin(0, 1, 1, 0)

	description = lipgloss.NewStyle().Margin(1, 0, 2, 0)

	archsTitle = title.Copy().Underline(true)
	archLabel  = lipgloss.NewStyle().
			MarginTop(1).
			MarginRight(1).
			Padding(0, 1).
			Background(archTagBg).
			Foreground(archTagFg)

	dockerPullCmdTitle = title.Copy().Underline(true).MarginTop(2)
	dockerPullCmdBox   = lipgloss.NewStyle().
				Background(dockerPullCmdBoxBg).
				Foreground(dockerPullCmdBoxFg).
				Padding(2).
				Margin(1, 0)
)
