package cli

import (
	"fmt"
	"strconv"
	"strings"

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
				Validate(func(value string) error {
					if strings.TrimSpace(value) == "" {
						return fmt.Errorf("project name cannot be empty")
					}

					return nil
				}),

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
				Validate(func(value string) error {
					count, err := strconv.Atoi(value)
					if err != nil {
						return fmt.Errorf("must be a number")
					}

					if count < 0 {
						return fmt.Errorf("cannot be negative")
					}

					return nil
				}),
		),
	)

	if err := form.Run(); err != nil {
		return project.CreateRequest{}, err
	}

	initialClips, err := strconv.Atoi(initialClipsInput)
	if err != nil {
		return project.CreateRequest{},
			fmt.Errorf("invalid clip count")
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
	var builder strings.Builder

	fmt.Fprintf(
		&builder,
		"Project: %s\n",
		plan.Root,
	)

	fmt.Fprintf(
		&builder,
		"Location: %s/%s\n\n",
		basePath,
		plan.Root,
	)

	builder.WriteString("Directories:\n")

	for _, directory := range plan.Directories {
		fmt.Fprintf(
			&builder,
			"  • %s\n",
			directory,
		)
	}

	var confirmed bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Create this project?").
				Description(builder.String()).
				Affirmative("Create").
				Negative("Cancel").
				Value(&confirmed),
		),
	)

	if err := form.Run(); err != nil {
		return false, err
	}

	return confirmed, nil
}

func init() {
	createCmd.AddCommand(createProjectCmd)
}
