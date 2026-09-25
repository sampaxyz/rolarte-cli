package prompt

import (
	"charm.land/huh/v2"
)

type ClipRequest struct {
	ProjectName string
	Count       int
}

func CreateClip(
	projects []string,
) (ClipRequest, error) {
	var selectedProject string

	countInput := "1"

	projectOptions := make(
		[]huh.Option[string],
		0,
		len(projects),
	)

	for _, projectName := range projects {
		projectOptions = append(
			projectOptions,
			huh.NewOption(
				projectName,
				projectName,
			),
		)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("System Project").
				Description(
					"Select the project that will receive the new clips.",
				).
				Options(projectOptions...).
				Value(&selectedProject),

			huh.NewInput().
				Title("Additional Clips").
				Description(
					"Number of new clips to create.",
				).
				Value(&countInput).
				Validate(validatePositiveInt),
		),
	)

	if err := form.Run(); err != nil {
		return ClipRequest{}, err
	}

	count, err := parseInt(countInput)
	if err != nil {
		return ClipRequest{}, err
	}

	return ClipRequest{
		ProjectName: selectedProject,
		Count:       count,
	}, nil
}
