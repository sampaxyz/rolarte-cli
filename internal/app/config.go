package app

import (
	"github.com/sampaxyz/rolarte-cli/internal/config"
)

func (app *App) LoadConfig() (config.Config, error) {
	return config.Load()
}

func (app *App) SaveConfig(cfg config.Config) error {
	return config.Save(cfg)
}

func (app *App) ConfigPath() (string, error) {
	return config.Path()
}
