package cli

import (
	"fmt"
	"path/filepath"

	"charm.land/huh/v2"
	"github.com/sampaxyz/rolarte-cli/internal/config"
	"github.com/sampaxyz/rolarte-cli/internal/filesystem"
	"github.com/sampaxyz/rolarte-cli/internal/project"
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create a new System Project",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf(
				"configuration is required before creating a project: %w",
				err,
			)
		}

		request, err := runCreateProjectForm()
		if err != nil {
			return err
		}

		plan, err := project.BuildPlan(request)
		if err != nil {
			return err
		}

		confirmed, err := confirmProjectPlan(
			cfg.SystemProjectsRoot,
			plan,
		)
		if err != nil {
			return err
		}

		if !confirmed {
			fmt.Println("Project creation cancelled.")
			return nil
		}

		executor := filesystem.NewExecutor()

		rootPath, err := executor.Execute(
			cfg.SystemProjectsRoot,
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
	},
}

func runCreateProjectForm() (project.CreateRequest, error) {
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

func confirmProjectPlan(
	basePath string,
	plan project.Plan,
) (bool, error) {
	description := fmt.Sprintf(
		"Project: %s\nLocation: %s",
		plan.Root,
		filepath.Join(basePath, plan.Root),
	)

	return confirmDirectories(
		"Create this project?",
		description,
		plan.Directories,
	)
}

func init() {
	createCmd.AddCommand(createProjectCmd)
}
