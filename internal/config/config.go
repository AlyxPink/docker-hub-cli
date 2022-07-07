package config

import (
	"strings"

	"github.com/docker/hack-docker-access-management-cli/internal/ui/components/section"
)

type Config struct {
	Sections []SectionConfig
}

type SectionConfig struct {
	Name    string
	Icon    string
	Section section.Model
}

func (v SectionConfig) Title() string {
	s := strings.Builder{}
	s.WriteString(v.Icon)
	s.WriteString(" ")
	s.WriteString(v.Name)
	return s.String()
}

func DefaultConfig() Config {
	return Config{
		Sections: []SectionConfig{
			{
				Name:    "Explore",
				Icon:    "",
				Section: section.NewModel(0),
			},
			{
				Name:    "My Repositories",
				Icon:    "",
				Section: section.NewModel(1),
			},
			{
				Name:    "My Organizations",
				Icon:    "",
				Section: section.NewModel(2),
			},
		},
	}
}
