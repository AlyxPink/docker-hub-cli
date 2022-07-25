package explore_sidebar

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/victorbersy/docker-hub-cli/internal/ui/components/sidebar"
)

var (
	archTagFg          = lipgloss.AdaptiveColor{Light: "#22223B", Dark: "#EDEDE9"}
	archTagBg          = lipgloss.AdaptiveColor{Light: "#BDE0FE", Dark: "#0077B6"}
	dockerPullCmdBoxBg = lipgloss.AdaptiveColor{Light: "#00F5D4", Dark: "#8AC926"}
	dockerPullCmdBoxFg = lipgloss.Color("#22223B")

	labelGlyph = sidebar.Title.Copy().Margin(0, 1, 1, 0)

	archsTitle = sidebar.Title.Copy().Underline(true)
	archLabel  = lipgloss.NewStyle().
			MarginTop(1).
			MarginRight(1).
			Padding(0, 1).
			Background(archTagBg).
			Foreground(archTagFg)

	dockerPullCmdTitle = sidebar.Title.Copy().Underline(true).MarginTop(2)
	dockerPullCmdBox   = lipgloss.NewStyle().
				Background(dockerPullCmdBoxBg).
				Foreground(dockerPullCmdBoxFg).
				Padding(2).
				Margin(1, 0)
)
