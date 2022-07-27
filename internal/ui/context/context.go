package context

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/victorbersy/docker-hub-cli/internal/config"
)

type ProgramContext struct {
	ScreenHeight      int
	ScreenWidth       int
	MainContentWidth  int
	MainContentHeight int
	Config            *config.Config
	View              config.ViewType
	Localizer         *i18n.Localizer
}

func (ctx *ProgramContext) GetViewsConfig() []config.ViewConfig {
	return ctx.Config.Views
}
