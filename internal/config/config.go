package config

import (
	"fmt"

	"github.com/AlyxPink/docker-hub-cli/internal/config/locales"
	"github.com/AlyxPink/docker-hub-cli/internal/ui/styles"
)

type ViewType string

const (
	ExploreView ViewType = "explore"
	MyReposView ViewType = "my_repos"
	MyOrgsView  ViewType = "my_orgs"
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
	localizer := locales.NewLocales()
	explore_title := localizer.L("tab_explore_title")
	my_repos_title := localizer.L("tab_my_repos_title")
	my_orgs_title := localizer.L("tab_my_orgs_title")
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
				Title: fmt.Sprint(styles.DefaultGlyphs.TabExplore, " ", explore_title),
				Type:  ExploreView,
			},
			{
				Title: fmt.Sprint(styles.DefaultGlyphs.TabMyRepos, "  ", my_repos_title),
				Type:  MyReposView,
			},
			{
				Title: fmt.Sprint(styles.DefaultGlyphs.TabMyOrgs, "  ", my_orgs_title),
				Type:  MyOrgsView,
			},
		},
	}
}
