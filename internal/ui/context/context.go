package context

import (
	"github.com/victorbersy/docker-hub-cli/internal/config"
)

type ProgramContext struct {
	ScreenHeight      int
	ScreenWidth       int
	MainContentWidth  int
	MainContentHeight int
	Config            *config.Config
	View              config.ViewType
}

func (ctx *ProgramContext) GetViewsConfig() []config.ViewConfig {
	return ctx.Config.Views
}
