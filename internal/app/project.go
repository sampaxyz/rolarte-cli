package app

import (
	"github.com/sampaxyz/rolarte-cli/internal/filesystem"
	"github.com/sampaxyz/rolarte-cli/internal/project"
)

type ProjectPlan struct {
	BasePath string
	Plan     project.Plan
}

func (app *App) PlanProject(
	request project.CreateRequest,
) (ProjectPlan, error) {
	cfg, err := app.LoadConfig()
	if err != nil {
		return ProjectPlan{}, err
	}

	plan, err := project.BuildPlan(request)
	if err != nil {
		return ProjectPlan{}, err
	}

	return ProjectPlan{
		BasePath: cfg.SystemProjectsRoot,
		Plan:     plan,
	}, nil
}

func (app *App) CreateProject(
	plan ProjectPlan,
) (string, error) {
	executor := filesystem.NewExecutor()

	return executor.Execute(
		plan.BasePath,
		plan.Plan,
	)
}
