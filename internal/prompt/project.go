package prompt

import (
	"charm.land/huh/v2"

	"github.com/sampaxyz/rolarte-cli/internal/project"
)

func CreateProject() (project.CreateRequest, error) {
	var projectName string
	var selectedSubprojects []project.Subproject

	initialClipsInput := "1"

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Description("Name of the new Rolarte System Project.").
				Value(&projectName).
				Validate(validateRequired),

			huh.NewMultiSelect[project.Subproject]().
				Title("Subprojects").
				Description("Select the subprojects to create.").
				Options(
					huh.NewOption(
						"Microfiction",
						project.SubprojectMicrofiction,
					).Selected(true),

					huh.NewOption(
						"System Overview",
						project.SubprojectSystemOverview,
					).Selected(true),

					huh.NewOption(
						"Podcast - Game Session",
						project.SubprojectPodcastGame,
					).Selected(true),

					huh.NewOption(
						"Podcast - Rolafter",
						project.SubprojectPodcastRolafter,
					).Selected(true),

					huh.NewOption(
						"Clips",
						project.SubprojectClips,
					).Selected(true),
				).
				Value(&selectedSubprojects),

			huh.NewInput().
				Title("Initial Clips").
				Description("Number of clips to create initially.").
				Value(&initialClipsInput).
				Validate(validatePositiveInt),
		),
	)

	if err := form.Run(); err != nil {
		return project.CreateRequest{}, err
	}

	initialClips, err := parseInt(initialClipsInput)
	if err != nil {
		return project.CreateRequest{}, err
	}

	return project.CreateRequest{
		Name:         projectName,
		Subprojects:  selectedSubprojects,
		InitialClips: initialClips,
	}, nil
}
