package tui

import tea "charm.land/bubbletea/v2"

type Action int

const (
	ActionCreateProject Action = iota
	ActionCreateClip
	ActionConfig
	ActionExit
)

type item struct {
	label  string
	action Action
}

type model struct {
	items []item

	cursor int

	projects []string
	rootPath string

	width  int
	height int

	selectedAction Action
}

func newModel(
	rootPath string,
	projects []string,
) model {
	return model{
		items: []item{
			{
				label:  "CREATE PROJECT",
				action: ActionCreateProject,
			},
			{
				label:  "CREATE CLIP",
				action: ActionCreateClip,
			},
			{
				label:  "CONFIGURATION",
				action: ActionConfig,
			},
			{
				label:  "EXIT",
				action: ActionExit,
			},
		},
		rootPath: rootPath,
		projects: projects,
	}
}

func (model model) Init() tea.Cmd {
	return nil
}
