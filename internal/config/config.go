package config

type ViewType string

const (
	ExploreView ViewType = "explore"
	MyReposView ViewType = "my_repos"
)

type ViewConfig struct {
	Title string
	Type  ViewType
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
	Views    []ViewConfig
	Defaults Defaults
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
		Views: []ViewConfig{
			{
				Title: " Explore",
				Type:  ExploreView,
			},
			{
				Title: "  My Repositories",
				Type:  MyReposView,
			},
		},
	}
}
