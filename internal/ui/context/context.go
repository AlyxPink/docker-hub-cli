package context

import "github.com/docker/hack-docker-access-management-cli/internal/config"

type ProgramContext struct {
	ScreenHeight      int
	ScreenWidth       int
	MainContentWidth  int
	MainContentHeight int
	Config            config.Config
}
