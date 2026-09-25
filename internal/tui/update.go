package tui

import (
	"fmt"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/sampaxyz/rolarte-cli/internal/project"
)

func (model model) Update(
	msg tea.Msg,
) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		model.width = msg.Width
		model.height = msg.Height

		return model, nil

	case dashboardLoadedMsg:
		return model.handleDashboardLoaded(msg)

	case projectPlannedMsg:
		if msg.err != nil {
			return model.fail(msg.err)
		}

		model.projectPlan = msg.plan
		model.projectPreviewPage = 0
		model.screen = screenProjectPreview

		return model, nil

	case projectCreatedMsg:
		if msg.err != nil {
			return model.fail(msg.err)
		}

		model.successTitle = "PROJECT CREATED"
		model.successMessage = fmt.Sprintf(
			"%s\n\n%s",
			model.projectPlan.Plan.Root,
			msg.rootPath,
		)

		model.screen = screenSuccess

		return model, nil

	case clipPlannedMsg:
		if msg.err != nil {
			return model.fail(msg.err)
		}

		model.clipPlan = msg.plan
		model.clipPreviewPage = 0
		model.screen = screenClipPreview

		return model, nil

	case clipsCreatedMsg:
		if msg.err != nil {
			return model.fail(msg.err)
		}

		model.successTitle = "CLIPS CREATED"
		model.successMessage = fmt.Sprintf(
			"%d clip(s) created in %s.",
			msg.count,
			msg.projectName,
		)

		model.screen = screenSuccess

		return model, nil

	case configSavedMsg:
		if msg.err != nil {
			return model.fail(msg.err)
		}

		model.rootPath = msg.rootPath

		model.successTitle = "CONFIGURATION SAVED"
		model.successMessage = msg.rootPath

		model.screen = screenSuccess

		return model, nil

	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return model, tea.Quit
		}

		return model.handleKey(msg)
	}

	return model, nil
}

func (model model) handleDashboardLoaded(
	msg dashboardLoadedMsg,
) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		return model.fail(msg.err)
	}

	if msg.notConfigured {
		model.screen = screenConfig
		model.rootPath = ""
		model.configRootInput.SetValue("")

		return model, model.configRootInput.Focus()
	}

	model.rootPath = msg.rootPath
	model.projects = msg.projects
	model.dashboardCursor = 0

	model.screen = screenDashboard

	return model, nil
}

func (model model) handleKey(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	switch model.screen {
	case screenDashboard:
		return model.updateDashboard(msg)

	case screenProjectName:
		return model.updateProjectName(msg)

	case screenProjectSubprojects:
		return model.updateProjectSubprojects(msg)

	case screenProjectClipCount:
		return model.updateProjectClipCount(msg)

	case screenProjectPreview:
		return model.updateProjectPreview(msg)

	case screenClipProject:
		return model.updateClipProject(msg)

	case screenClipCount:
		return model.updateClipCount(msg)

	case screenClipPreview:
		return model.updateClipPreview(msg)

	case screenConfig:
		return model.updateConfig(msg)

	case screenSuccess:
		return model.updateSuccess(msg)

	case screenError:
		return model.updateError(msg)
	}

	return model, nil
}

func (model model) updateDashboard(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if model.dashboardCursor > 0 {
			model.dashboardCursor--
		}

	case "down", "j":
		if model.dashboardCursor <
			len(model.dashboardItems)-1 {
			model.dashboardCursor++
		}

	case "enter":
		item := model.dashboardItems[model.dashboardCursor]

		switch item.action {
		case actionCreateProject:
			return model,
				model.startCreateProject()

		case actionCreateClip:
			return model,
				model.startCreateClip()

		case actionConfig:
			return model,
				model.startConfig()

		case actionExit:
			return model, tea.Quit
		}

	case "q":
		return model, tea.Quit
	}

	return model, nil
}

func (model model) updateProjectName(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		model.projectNameInput.Blur()
		model.screen = screenDashboard
		model.validationError = ""

		return model, nil

	case "enter":
		name := strings.TrimSpace(
			model.projectNameInput.Value(),
		)

		if name == "" {
			model.validationError =
				"Project name cannot be empty."

			return model, nil
		}

		model.projectNameInput.Blur()
		model.validationError = ""
		model.screen = screenProjectSubprojects

		return model, nil
	}

	var cmd tea.Cmd

	model.projectNameInput, cmd =
		model.projectNameInput.Update(msg)

	model.validationError = ""

	return model, cmd
}

func (model model) updateProjectSubprojects(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	if msg.Code == tea.KeySpace {
		index := model.projectSubprojectCursor

		model.projectSelected[index] =
			!model.projectSelected[index]

		model.validationError = ""

		return model, nil
	}

	switch msg.String() {
	case "up", "k":
		if model.projectSubprojectCursor > 0 {
			model.projectSubprojectCursor--
		}

	case "down", "j":
		if model.projectSubprojectCursor <
			len(projectOptions)-1 {
			model.projectSubprojectCursor++
		}

	case "b", "esc":
		model.screen = screenProjectName

		return model,
			model.projectNameInput.Focus()

	case "enter":
		if len(model.selectedSubprojects()) == 0 {
			model.validationError =
				"Select at least one subproject."

			return model, nil
		}

		model.validationError = ""

		if model.clipsSelected() {
			model.screen = screenProjectClipCount

			return model,
				model.projectClipCountInput.Focus()
		}

		request := project.CreateRequest{
			Name: model.projectNameInput.Value(),

			Subprojects: model.selectedSubprojects(),

			InitialClips: 0,
		}

		return model,
			planProjectCmd(
				model.application,
				request,
			)
	}

	return model, nil
}

func (model model) updateProjectClipCount(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		model.projectClipCountInput.Blur()
		model.screen = screenProjectSubprojects

		return model, nil

	case "enter":
		count, err := parsePositiveInt(
			model.projectClipCountInput.Value(),
		)

		if err != nil {
			model.validationError = err.Error()

			return model, nil
		}

		model.projectClipCountInput.Blur()
		model.validationError = ""

		request := project.CreateRequest{
			Name: model.projectNameInput.Value(),

			Subprojects: model.selectedSubprojects(),

			InitialClips: count,
		}

		return model,
			planProjectCmd(
				model.application,
				request,
			)
	}

	var cmd tea.Cmd

	model.projectClipCountInput, cmd =
		model.projectClipCountInput.Update(msg)

	model.validationError = ""

	return model, cmd
}

func (model model) updateProjectPreview(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	totalPages := pageCount(
		len(model.projectPlan.Plan.Directories),
		previewPageSize,
	)

	switch msg.String() {
	case "left", "h":
		if model.projectPreviewPage > 0 {
			model.projectPreviewPage--
		}

	case "right", "l":
		if model.projectPreviewPage < totalPages-1 {
			model.projectPreviewPage++
		}

	case "enter":
		return model,
			createProjectCmd(
				model.application,
				model.projectPlan,
			)

	case "b", "esc":
		if model.clipsSelected() {
			model.screen = screenProjectClipCount

			return model,
				model.projectClipCountInput.Focus()
		}

		model.screen = screenProjectSubprojects
	}

	return model, nil
}

func (model model) updateClipProject(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if model.clipProjectCursor > 0 {
			model.clipProjectCursor--
		}

	case "down", "j":
		if model.clipProjectCursor <
			len(model.projects)-1 {
			model.clipProjectCursor++
		}

	case "enter":
		model.clipCountInput.SetValue("1")
		model.screen = screenClipCount

		return model,
			model.clipCountInput.Focus()

	case "esc":
		model.screen = screenDashboard
	}

	return model, nil
}

func (model model) updateClipCount(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		model.clipCountInput.Blur()
		model.screen = screenClipProject

		return model, nil

	case "enter":
		count, err := parsePositiveInt(
			model.clipCountInput.Value(),
		)

		if err != nil {
			model.validationError = err.Error()

			return model, nil
		}

		projectName :=
			model.projects[model.clipProjectCursor]

		model.clipCountInput.Blur()
		model.validationError = ""

		return model,
			planClipsCmd(
				model.application,
				projectName,
				count,
			)
	}

	var cmd tea.Cmd

	model.clipCountInput, cmd =
		model.clipCountInput.Update(msg)

	model.validationError = ""

	return model, cmd
}

func (model model) updateClipPreview(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	totalPages := pageCount(
		len(model.clipPlan.Plan.Directories),
		previewPageSize,
	)

	switch msg.String() {
	case "left", "h":
		if model.clipPreviewPage > 0 {
			model.clipPreviewPage--
		}

	case "right", "l":
		if model.clipPreviewPage < totalPages-1 {
			model.clipPreviewPage++
		}

	case "enter":
		count, err := parsePositiveInt(
			model.clipCountInput.Value(),
		)

		if err != nil {
			return model.fail(err)
		}

		return model,
			createClipsCmd(
				model.application,
				model.clipPlan,
				count,
			)

	case "b", "esc":
		model.screen = screenClipCount

		return model,
			model.clipCountInput.Focus()
	}

	return model, nil
}

func (model model) updateConfig(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		if model.rootPath == "" {
			return model, tea.Quit
		}

		model.configRootInput.Blur()
		model.screen = screenDashboard

		return model, nil

	case "enter":
		rootPath := strings.TrimSpace(
			model.configRootInput.Value(),
		)

		if rootPath == "" {
			model.validationError =
				"System Projects Root cannot be empty."

			return model, nil
		}

		model.configRootInput.Blur()
		model.validationError = ""

		return model,
			saveConfigCmd(
				model.application,
				rootPath,
			)
	}

	var cmd tea.Cmd

	model.configRootInput, cmd =
		model.configRootInput.Update(msg)

	model.validationError = ""

	return model, cmd
}

func (model model) updateSuccess(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "esc":
		model.screen = screenLoading

		return model,
			loadDashboardCmd(
				model.application,
			)
	}

	return model, nil
}

func (model model) updateError(
	msg tea.KeyPressMsg,
) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "esc", "b":
		model.currentError = nil
		model.screen = screenLoading

		return model,
			loadDashboardCmd(
				model.application,
			)

	case "q":
		return model, tea.Quit
	}

	return model, nil
}

func (model model) fail(
	err error,
) (tea.Model, tea.Cmd) {
	model.currentError = err
	model.screen = screenError

	return model, nil
}

func parsePositiveInt(
	value string,
) (int, error) {
	count, err := strconv.Atoi(
		strings.TrimSpace(value),
	)

	if err != nil {
		return 0, fmt.Errorf(
			"Enter a valid number.",
		)
	}

	if count < 1 {
		return 0, fmt.Errorf(
			"Value must be greater than zero.",
		)
	}

	return count, nil
}
