package tui

import (
	"errors"

	"github.com/sampaxyz/rolarte-cli/internal/app"
	"github.com/sampaxyz/rolarte-cli/internal/config"
	"github.com/sampaxyz/rolarte-cli/internal/interactive"
)

func Run(
	application *app.App,
) error {
	for {
		cfg, err := application.LoadConfig()

		if errors.Is(
			err,
			config.ErrSystemProjectsRootNotConfigured,
		) {
			if err := interactive.Configure(application); err != nil {
				return err
			}

			continue
		}

		if err != nil {
			return err
		}

		projects, err := application.ListProjects(
			cfg.SystemProjectsRoot,
		)
		if err != nil {
			return err
		}

		action, err := RunDashboard(
			cfg.SystemProjectsRoot,
			projects,
		)
		if err != nil {
			return err
		}

		switch action {
		case ActionCreateProject:
			if err := interactive.CreateProject(
				application,
			); err != nil {
				return err
			}

		case ActionCreateClip:
			if err := interactive.CreateClip(
				application,
			); err != nil {
				return err
			}

		case ActionConfig:
			if err := interactive.Configure(
				application,
			); err != nil {
				return err
			}

		case ActionExit:
			return nil
		}
	}
}
