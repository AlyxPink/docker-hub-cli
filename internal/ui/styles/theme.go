package styles

import "github.com/charmbracelet/lipgloss"

var indigo = lipgloss.AdaptiveColor{Light: "#FFAFCC", Dark: "#383B5B"}

type Theme struct {
	MainText      lipgloss.AdaptiveColor
	SubleMainText lipgloss.AdaptiveColor

	Border          lipgloss.AdaptiveColor
	SecondaryBorder lipgloss.AdaptiveColor

	SuccessText lipgloss.AdaptiveColor
	ErrorText   lipgloss.AdaptiveColor

	FaintBorder lipgloss.AdaptiveColor
	FaintText   lipgloss.AdaptiveColor

	SelectedBackground lipgloss.AdaptiveColor
	SecondaryText      lipgloss.AdaptiveColor

	ArchTagBg              lipgloss.AdaptiveColor
	ArchTagFg              lipgloss.AdaptiveColor
	DockerPullCmdBoxBg     lipgloss.AdaptiveColor
	DockerPullCmdBoxFg     lipgloss.AdaptiveColor
	LabelCommunity         lipgloss.AdaptiveColor
	LabelDockerOfficial    lipgloss.AdaptiveColor
	LabelOpenSourceProgram lipgloss.AdaptiveColor
	LabelVerifiedPublisher lipgloss.AdaptiveColor
	StatsDownloads         lipgloss.AdaptiveColor
	StatsStars             lipgloss.AdaptiveColor
}

type Glyphs struct {
	LabelCommunity         string
	LabelDockerOfficial    string
	LabelOpenSourceProgram string
	LabelVerifiedPublisher string

	IsPrivate string
	Private   string
	Public    string

	StatsDownloads string
	StatsStars     string
}

var subtleIndigo = lipgloss.AdaptiveColor{Light: "#FFC8DD", Dark: "#242347"}

var DefaultTheme = Theme{
	MainText:      lipgloss.AdaptiveColor{Light: "#242347", Dark: "#E2E1ED"},
	SubleMainText: subtleIndigo,

	Border:          lipgloss.AdaptiveColor{Light: indigo.Light, Dark: indigo.Dark},
	SecondaryBorder: lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#39386B"},

	SuccessText: lipgloss.AdaptiveColor{Light: "#3DF294", Dark: "#06D6A0"},
	ErrorText:   lipgloss.AdaptiveColor{Light: "#F23D5C", Dark: "#EF476F"},

	FaintBorder: lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#2B2B40"},
	FaintText:   lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#3E4057"},

	SelectedBackground: lipgloss.AdaptiveColor{Light: subtleIndigo.Light, Dark: "#39386B"},
	SecondaryText:      lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#666CA6"},

	ArchTagBg:              lipgloss.AdaptiveColor{Light: "#BDE0FE", Dark: "#0077B6"},
	ArchTagFg:              lipgloss.AdaptiveColor{Light: "#22223B", Dark: "#EDEDE9"},
	DockerPullCmdBoxBg:     lipgloss.AdaptiveColor{Light: "#00F5D4", Dark: "#8AC926"},
	DockerPullCmdBoxFg:     lipgloss.AdaptiveColor{Light: "#22223B", Dark: "#22223B"},
	LabelCommunity:         lipgloss.AdaptiveColor{Light: "#BDE0FE", Dark: "#BDE0FE"},
	LabelDockerOfficial:    lipgloss.AdaptiveColor{Light: "#83C5BE", Dark: "#2E7F74"},
	LabelOpenSourceProgram: lipgloss.AdaptiveColor{Light: "#7D2EFF", Dark: "#7D2EFF"},
	LabelVerifiedPublisher: lipgloss.AdaptiveColor{Light: "#48CAE4", Dark: "#086DD7"},
	StatsDownloads:         lipgloss.AdaptiveColor{Light: "#00BBF9", Dark: "#00BBF9"},
	StatsStars:             lipgloss.AdaptiveColor{Light: "#FFB703", Dark: "#FFB703"},
}

var DefaultGlyphs = Glyphs{
	LabelCommunity:         "וֹ",
	LabelDockerOfficial:    "",
	LabelOpenSourceProgram: "",
	LabelVerifiedPublisher: "﫠",

	IsPrivate: "",
	Private:   "",
	Public:    "",

	StatsDownloads: "",
	StatsStars:     "",
}

var (
	SingleRuneWidth    = 4
	MainContentPadding = 1
)
