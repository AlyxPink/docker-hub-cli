package config

import (
	"strings"

	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/view"
)

type Config struct {
	Views []ViewConfig
}

type ViewConfig struct {
	Name string
	Icon string
	View view.Model
}

func (v ViewConfig) Title() string {
	s := strings.Builder{}
	s.WriteString(v.Icon)
	s.WriteString(" ")
	s.WriteString(v.Name)
	return s.String()
}

func DefaultConfig() Config {
	return Config{
		Views: []ViewConfig{
			{
				Name: "Explore",
				Icon: "",
				View: view.NewModel(0),
			},
			{
				Name: "My Repositories",
				Icon: "",
				View: view.NewModel(1),
			},
			{
				Name: "My Organizations",
				Icon: "",
				View: view.NewModel(2),
			},
		},
	}
}
