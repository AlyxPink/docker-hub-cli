package styles

type Glyphs struct {
	TabExplore string
	TabMyRepos string

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

var DefaultGlyphs = Glyphs{
	TabExplore: "",
	TabMyRepos: "",

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
