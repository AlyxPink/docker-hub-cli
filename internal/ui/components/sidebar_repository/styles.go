package sidebar_repository

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	labelDockerOfficial    = lipgloss.Color("#2E7F74")
	labelVerifiedPublisher = lipgloss.Color("#086DD7")
	labelOpenSourceProgram = lipgloss.Color("#7D2EFF")
	archTagFg              = lipgloss.AdaptiveColor{Light: "#22223b", Dark: "#EDEDE9"}
	archTagBg              = lipgloss.AdaptiveColor{Light: "#bde0fe", Dark: "#0077b6"}
	dockerImageTagFg       = lipgloss.Color("#22223b")
	dockerImageTagBg       = lipgloss.Color("#EDEDE9")
	dockerPullCmdBoxBg     = lipgloss.AdaptiveColor{Light: "#00f5d4", Dark: "#8ac926"}
	dockerPullCmdBoxFg     = lipgloss.Color("#22223b")

	label = lipgloss.NewStyle().Copy().
		Padding(0, 1).
		Bold(false)

	labelTitle = label.Copy().
			Underline(true).
			Bold(true).
			Margin(0, 1, 1, 0)

	archTitle = labelTitle
	archTag   = label.Copy().
			Background(archTagBg).
			Foreground(archTagFg).
			Margin(0, 1)

	dockerImageTitle = labelTitle.Copy()
	dockerImageTag   = label.Copy().
				Background(dockerImageTagBg).
				Foreground(dockerImageTagFg).
				Margin(0, 1)

	dockerPullCmdTitle = labelTitle.Copy()
	dockerPullCmdBox   = label.Copy().
				Background(dockerPullCmdBoxBg).
				Foreground(dockerPullCmdBoxFg).
				Padding(2).
				Margin(2)
)
