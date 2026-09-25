package interactive

import (
	"fmt"
	"path/filepath"

	"github.com/sampaxyz/rolarte-cli/internal/app"
	"github.com/sampaxyz/rolarte-cli/internal/prompt"
)

func CreateProject(
	application *app.App,
) error {
	request, err := prompt.CreateProject()
	if err != nil {
		return err
	}

	plan, err := application.PlanProject(
		request,
	)
	if err != nil {
		return err
	}

	description := fmt.Sprintf(
		"Project: %s\nLocation: %s",
		plan.Plan.Root,
		filepath.Join(
			plan.BasePath,
			plan.Plan.Root,
		),
	)

	confirmed, err := prompt.ConfirmDirectories(
		"Create this project?",
		description,
		plan.Plan.Directories,
	)
	if err != nil {
		return err
	}

	if !confirmed {
		return nil
	}

	rootPath, err := application.CreateProject(
		plan,
	)
	if err != nil {
		return err
	}

	fmt.Printf(
		"\nProject created successfully:\n%s\n",
		rootPath,
	)

	return nil
}
