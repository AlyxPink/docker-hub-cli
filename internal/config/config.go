package config

type ViewType string

const (
	ExploreView ViewType = "explore"
)

type SectionConfig struct {
	Title string
	Limit *int
}

type PreviewConfig struct {
	Open  bool
	Width int
}

type Defaults struct {
	Preview PreviewConfig
	View    ViewType
}

type Keybinding struct {
	Key     string
	Command string
}

type Config struct {
	ExploreSections []SectionConfig
	Defaults        Defaults
}

func GetDefaultConfig() Config {
	return Config{
		Defaults: Defaults{
			Preview: PreviewConfig{
				Open:  false,
				Width: 50,
			},
			View: ExploreView,
		},
		ExploreSections: []SectionConfig{
			{
				Title: " Explore",
			},
			{
				Title: " My Repositories",
			},
			{
				Title: " My Organizations",
			},
		},
	}
}
