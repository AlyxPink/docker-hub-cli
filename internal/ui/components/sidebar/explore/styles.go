package explore_sidebar

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/components/sidebar"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/styles"
)

var (
	labelGlyph = sidebar.Title.Copy().Margin(0, 1, 1, 0)

	archsTitle = sidebar.Title.Copy().Underline(true)
	archLabel  = lipgloss.NewStyle().
			MarginTop(1).
			MarginRight(1).
			Padding(0, 1).
			Background(styles.DefaultTheme.ArchTagBg).
			Foreground(styles.DefaultTheme.ArchTagFg)

	dockerPullCmdTitle = sidebar.Title.Copy().Underline(true).MarginTop(2)
	dockerPullCmdBox   = lipgloss.NewStyle().
				Background(styles.DefaultTheme.DockerPullCmdBoxBg).
				Foreground(styles.DefaultTheme.DockerPullCmdBoxFg).
				Padding(2).
				Margin(1, 0)
)
