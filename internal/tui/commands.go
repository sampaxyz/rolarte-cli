package tui

import (
	"errors"

	tea "charm.land/bubbletea/v2"

	"github.com/sampaxyz/rolarte-cli/internal/app"
	"github.com/sampaxyz/rolarte-cli/internal/config"
	"github.com/sampaxyz/rolarte-cli/internal/project"
)

func loadDashboardCmd(
	application *app.App,
) tea.Cmd {
	return func() tea.Msg {
		cfg, err := application.LoadConfig()

		if errors.Is(
			err,
			config.ErrSystemProjectsRootNotConfigured,
		) {
			return dashboardLoadedMsg{
				notConfigured: true,
			}
		}

		if err != nil {
			return dashboardLoadedMsg{
				err: err,
			}
		}

		projects, err := application.ListProjects(
			cfg.SystemProjectsRoot,
		)

		return dashboardLoadedMsg{
			rootPath: cfg.SystemProjectsRoot,
			projects: projects,
			err:      err,
		}
	}
}

func planProjectCmd(
	application *app.App,
	request project.CreateRequest,
) tea.Cmd {
	return func() tea.Msg {
		plan, err := application.PlanProject(
			request,
		)

		return projectPlannedMsg{
			plan: plan,
			err:  err,
		}
	}
}

func createProjectCmd(
	application *app.App,
	plan app.ProjectPlan,
) tea.Cmd {
	return func() tea.Msg {
		rootPath, err := application.CreateProject(
			plan,
		)

		return projectCreatedMsg{
			rootPath: rootPath,
			err:      err,
		}
	}
}

func planClipsCmd(
	application *app.App,
	projectName string,
	count int,
) tea.Cmd {
	return func() tea.Msg {
		plan, err := application.PlanClips(
			projectName,
			count,
		)

		return clipPlannedMsg{
			plan: plan,
			err:  err,
		}
	}
}

func createClipsCmd(
	application *app.App,
	plan app.ClipPlan,
	count int,
) tea.Cmd {
	return func() tea.Msg {
		err := application.CreateClips(
			plan,
		)

		return clipsCreatedMsg{
			count:       count,
			projectName: plan.ProjectName,
			err:         err,
		}
	}
}

func saveConfigCmd(
	application *app.App,
	rootPath string,
) tea.Cmd {
	return func() tea.Msg {
		cfg := config.Config{
			SystemProjectsRoot: rootPath,
		}

		err := application.SaveConfig(cfg)

		return configSavedMsg{
			rootPath: rootPath,
			err:      err,
		}
	}
}
