package tui

import (
	tea "charm.land/bubbletea/v2"

	"github.com/sampaxyz/rolarte-cli/internal/app"
)

func Run(application *app.App) error {
	program := tea.NewProgram(
		newModel(application),
	)

	_, err := program.Run()

	return err
}
