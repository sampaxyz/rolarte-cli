package tui

import tea "charm.land/bubbletea/v2"

func RunDashboard(
	rootPath string,
	projects []string,
) (Action, error) {
	initialModel := newModel(
		rootPath,
		projects,
	)

	program := tea.NewProgram(
		initialModel,
	)

	finalModel, err := program.Run()
	if err != nil {
		return ActionExit, err
	}

	final := finalModel.(model)

	return final.selectedAction, nil
}
