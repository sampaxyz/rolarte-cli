package app

import (
	"github.com/sampaxyz/rolarte-cli/internal/filesystem"
)

func (app *App) ListProjects(
	rootPath string,
) ([]string, error) {
	return filesystem.ListProjects(rootPath)
}
