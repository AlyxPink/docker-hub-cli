package config

type ViewType string

const (
	PRsView ViewType = "prs"
)

type SectionConfig struct {
	Title string
	Limit *int `yaml:"limit,omitempty"`
}

type PreviewConfig struct {
	Open  bool
	Width int
}

type Defaults struct {
	Preview  PreviewConfig `yaml:"preview"`
	PrsLimit int           `yaml:"prsLimit"`
	View     ViewType      `yaml:"view"`
}

type Keybinding struct {
	Key     string `yaml:"key"`
	Command string `yaml:"command"`
}

type Keybindings struct {
	Prs []Keybinding `yaml:"prs"`
}

type Config struct {
	PRSections  []SectionConfig   `yaml:"prSections"`
	Defaults    Defaults          `yaml:"defaults"`
	Keybindings Keybindings       `yaml:"keybindings"`
	RepoPaths   map[string]string `yaml:"repoPaths"`
}

func GetDefaultConfig() Config {
	return Config{
		Defaults: Defaults{
			Preview: PreviewConfig{
				Open:  false,
				Width: 50,
			},
			PrsLimit: 20,
			View:     PRsView,
		},
		PRSections: []SectionConfig{
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
		Keybindings: Keybindings{
			Prs: []Keybinding{},
		},
		RepoPaths: map[string]string{},
	}
}
