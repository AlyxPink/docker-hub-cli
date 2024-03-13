package context

import (
	"github.com/AlyxPink/docker-hub-cli/internal/config"
	"github.com/AlyxPink/docker-hub-cli/internal/config/locales"
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
