package tui

import (
	"github.com/sampaxyz/rolarte-cli/internal/app"
)

type dashboardLoadedMsg struct {
	rootPath      string
	projects      []string
	notConfigured bool
	err           error
}

type projectPlannedMsg struct {
	plan app.ProjectPlan
	err  error
}

type projectCreatedMsg struct {
	rootPath string
	err      error
}

type clipPlannedMsg struct {
	plan app.ClipPlan
	err  error
}

type clipsCreatedMsg struct {
	count       int
	projectName string
	err         error
}

type configSavedMsg struct {
	rootPath string
	err      error
}
