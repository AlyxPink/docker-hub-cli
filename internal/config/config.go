package config

import "strings"

type SectionConfig struct {
	Name string
	Icon string
}

type Config struct {
	Sections []SectionConfig
}

func (sc SectionConfig) Title() string {
	s := strings.Builder{}
	// TODO: Remove icon if not supported (requires Nerd font)
	s.WriteString(sc.Icon)
	s.WriteString(" ")
	s.WriteString(sc.Name)
	return s.String()
}

func DefaultConfig() Config {
	return Config{
		Sections: []SectionConfig{
			{
				Name: "Explore",
				Icon: "",
			},
			{
				Name: "My Repositories",
				Icon: "",
			},
			{
				Name: "My Organizations",
				Icon: "",
			},
		},
	}
}
