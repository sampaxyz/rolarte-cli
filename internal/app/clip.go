package app

import (
	"path/filepath"

	"github.com/sampaxyz/rolarte-cli/internal/filesystem"
	"github.com/sampaxyz/rolarte-cli/internal/project"
)

type ClipPlan struct {
	ProjectName string
	ProjectPath string
	Plan        project.ClipPlan
}

func (app *App) PlanClips(
	projectName string,
	count int,
) (ClipPlan, error) {
	cfg, err := app.LoadConfig()
	if err != nil {
		return ClipPlan{}, err
	}

	projectPath := filepath.Join(
		cfg.SystemProjectsRoot,
		projectName,
	)

	nextIndex, err := filesystem.NextClipIndex(
		projectPath,
	)
	if err != nil {
		return ClipPlan{}, err
	}

	plan, err := project.BuildClipPlan(
		nextIndex,
		count,
	)
	if err != nil {
		return ClipPlan{}, err
	}

	return ClipPlan{
		ProjectName: projectName,
		ProjectPath: projectPath,
		Plan:        plan,
	}, nil
}

func (app *App) CreateClips(
	plan ClipPlan,
) error {
	executor := filesystem.NewExecutor()

	return executor.ExecuteDirectories(
		plan.ProjectPath,
		plan.Plan.Directories,
	)
}
