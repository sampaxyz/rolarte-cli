package tui

import (
	"errors"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"

	"github.com/sampaxyz/rolarte-cli/internal/app"
	"github.com/sampaxyz/rolarte-cli/internal/project"
)

const previewPageSize = 10

var errNoProjects = errors.New(
	"there are no System Projects yet",
)

type screen int

const (
	screenLoading screen = iota

	screenDashboard

	screenProjectName
	screenProjectSubprojects
	screenProjectClipCount
	screenProjectPreview

	screenClipProject
	screenClipCount
	screenClipPreview

	screenConfig

	screenSuccess
	screenError
)

type dashboardAction int

const (
	actionCreateProject dashboardAction = iota
	actionCreateClip
	actionConfig
	actionExit
)

type dashboardItem struct {
	label  string
	action dashboardAction
}

type projectOption struct {
	label string
	value project.Subproject
}

var projectOptions = [...]projectOption{
	{
		label: "Microfiction",
		value: project.SubprojectMicrofiction,
	},
	{
		label: "System Overview",
		value: project.SubprojectSystemOverview,
	},
	{
		label: "Podcast - Game Session",
		value: project.SubprojectPodcastGame,
	},
	{
		label: "Podcast - Rolafter",
		value: project.SubprojectPodcastRolafter,
	},
	{
		label: "Clips",
		value: project.SubprojectClips,
	},
}

type model struct {
	application *app.App

	screen screen

	width  int
	height int

	rootPath string
	projects []string

	// Dashboard
	dashboardItems  []dashboardItem
	dashboardCursor int

	// Create Project
	projectNameInput        textinput.Model
	projectSubprojectCursor int
	projectSelected         [len(projectOptions)]bool
	projectClipCountInput   textinput.Model
	projectPlan             app.ProjectPlan
	projectPreviewPage      int

	// Create Clip
	clipProjectCursor int
	clipCountInput    textinput.Model
	clipPlan          app.ClipPlan
	clipPreviewPage   int

	// Config
	configRootInput textinput.Model

	// Generic feedback
	validationError string
	successTitle    string
	successMessage  string
	currentError    error
}

func newModel(
	application *app.App,
) model {
	projectNameInput := textinput.New()
	projectNameInput.Placeholder = "Outta Bubblegum"
	projectNameInput.SetWidth(50)

	projectClipCountInput := textinput.New()
	projectClipCountInput.SetValue("1")
	projectClipCountInput.SetWidth(10)

	clipCountInput := textinput.New()
	clipCountInput.SetValue("1")
	clipCountInput.SetWidth(10)

	configRootInput := textinput.New()
	configRootInput.Placeholder = "/Users/.../Rolarte/Systems"
	configRootInput.SetWidth(60)

	return model{
		application: application,

		screen: screenLoading,

		dashboardItems: []dashboardItem{
			{
				label:  "CREATE PROJECT",
				action: actionCreateProject,
			},
			{
				label:  "CREATE CLIP",
				action: actionCreateClip,
			},
			{
				label:  "CONFIGURATION",
				action: actionConfig,
			},
			{
				label:  "EXIT",
				action: actionExit,
			},
		},

		projectNameInput:      projectNameInput,
		projectClipCountInput: projectClipCountInput,
		clipCountInput:        clipCountInput,
		configRootInput:       configRootInput,

		projectSelected: [len(projectOptions)]bool{
			true,
			true,
			true,
			true,
			true,
		},
	}
}

func (model model) Init() tea.Cmd {
	return loadDashboardCmd(
		model.application,
	)
}

func (model *model) startCreateProject() tea.Cmd {
	model.screen = screenProjectName

	model.validationError = ""

	model.projectNameInput.SetValue("")
	model.projectClipCountInput.SetValue("1")

	model.projectSubprojectCursor = 0
	model.projectPreviewPage = 0

	model.projectSelected = [len(projectOptions)]bool{
		true,
		true,
		true,
		true,
		true,
	}

	return model.projectNameInput.Focus()
}

func (model *model) startCreateClip() tea.Cmd {
	model.validationError = ""
	model.clipPreviewPage = 0

	if len(model.projects) == 0 {
		model.currentError = errNoProjects
		model.screen = screenError
		return nil
	}

	model.clipProjectCursor = 0
	model.clipCountInput.SetValue("1")

	model.screen = screenClipProject

	return nil
}

func (model *model) startConfig() tea.Cmd {
	model.validationError = ""

	model.configRootInput.SetValue(
		model.rootPath,
	)

	model.screen = screenConfig

	return model.configRootInput.Focus()
}

func (model *model) selectedSubprojects() []project.Subproject {
	selected := make(
		[]project.Subproject,
		0,
		len(projectOptions),
	)

	for index, option := range projectOptions {
		if model.projectSelected[index] {
			selected = append(
				selected,
				option.value,
			)
		}
	}

	return selected
}

func (model model) clipsSelected() bool {
	for index, option := range projectOptions {
		if option.value == project.SubprojectClips {
			return model.projectSelected[index]
		}
	}

	return false
}

func pageCount(
	total int,
	pageSize int,
) int {
	if total == 0 {
		return 1
	}

	return (total + pageSize - 1) / pageSize
}

func pageSlice(
	items []string,
	page int,
	pageSize int,
) []string {
	start := page * pageSize

	if start >= len(items) {
		return nil
	}

	end := start + pageSize

	if end > len(items) {
		end = len(items)
	}

	return items[start:end]
}
