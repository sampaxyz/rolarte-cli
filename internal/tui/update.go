package tui

import tea "charm.land/bubbletea/v2"

func (model model) Update(
	msg tea.Msg,
) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		model.width = msg.Width
		model.height = msg.Height

	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if model.cursor > 0 {
				model.cursor--
			}

		case "down", "j":
			if model.cursor < len(model.items)-1 {
				model.cursor++
			}

		case "enter":
			model.selectedAction =
				model.items[model.cursor].action

			return model, tea.Quit

		case "q", "ctrl+c":
			model.selectedAction = ActionExit

			return model, tea.Quit
		}
	}

	return model, nil
}
