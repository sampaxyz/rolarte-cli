package interactive

import (
	"fmt"

	"github.com/sampaxyz/rolarte-cli/internal/app"
	"github.com/sampaxyz/rolarte-cli/internal/prompt"
)

func CreateClip(
	application *app.App,
) error {
	cfg, err := application.LoadConfig()
	if err != nil {
		return err
	}

	projects, err := application.ListProjects(
		cfg.SystemProjectsRoot,
	)
	if err != nil {
		return err
	}

	if len(projects) == 0 {
		return fmt.Errorf(
			"no System Projects found in %s",
			cfg.SystemProjectsRoot,
		)
	}

	request, err := prompt.CreateClip(
		projects,
	)
	if err != nil {
		return err
	}

	plan, err := application.PlanClips(
		request.ProjectName,
		request.Count,
	)
	if err != nil {
		return err
	}

	description := fmt.Sprintf(
		"Project: %s\nLocation: %s",
		plan.ProjectName,
		plan.ProjectPath,
	)

	confirmed, err := prompt.ConfirmDirectories(
		"Create these clips?",
		description,
		plan.Plan.Directories,
	)
	if err != nil {
		return err
	}

	if !confirmed {
		return nil
	}

	if err := application.CreateClips(plan); err != nil {
		return err
	}

	fmt.Printf(
		"\nCreated %d clip(s) in %s.\n",
		request.Count,
		request.ProjectName,
	)

	return nil
}
