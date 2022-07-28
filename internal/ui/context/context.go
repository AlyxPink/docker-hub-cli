package context

import (
	"github.com/victorbersy/docker-hub-cli/internal/config"
	"github.com/victorbersy/docker-hub-cli/internal/config/locales"
)

type ProgramContext struct {
	ScreenHeight      int
	ScreenWidth       int
	MainContentWidth  int
	MainContentHeight int
	Config            *config.Config
	View              config.ViewType
	Localizer         locales.Locales
}

func (ctx *ProgramContext) GetViewsConfig() []config.ViewConfig {
	return ctx.Config.Views
}
