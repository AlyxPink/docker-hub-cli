package context

import "github.com/VictorBersy/docker-hub-cli/internal/config"

type ProgramContext struct {
	ScreenHeight      int
	ScreenWidth       int
	MainContentWidth  int
	MainContentHeight int
	Config            *config.Config
	View              config.ViewType
}

func (ctx *ProgramContext) GetViewSectionsConfig() []config.SectionConfig {
	return ctx.Config.ExploreSections
}
