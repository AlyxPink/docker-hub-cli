package config

type ViewType string

const (
	ExploreView ViewType = "explore"
)

type ViewConfig struct {
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
	ExploreViews []ViewConfig
	Defaults     Defaults
}

func GetDefaultConfig() Config {
	return Config{
		Defaults: Defaults{
			Preview: PreviewConfig{
				Open:  true,
				Width: 70,
			},
			View: ExploreView,
		},
		ExploreViews: []ViewConfig{
			{
				Title: " Explore",
			},
			{
				Title: "  My Repositories",
			},
			{
				Title: "  My Organizations",
			},
		},
	}
}
