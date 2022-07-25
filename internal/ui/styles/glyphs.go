package styles

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
