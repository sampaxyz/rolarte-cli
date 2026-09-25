package cli

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"charm.land/huh/v2"
	"github.com/sampaxyz/rolarte-cli/internal/config"
	"github.com/sampaxyz/rolarte-cli/internal/filesystem"
	"github.com/sampaxyz/rolarte-cli/internal/project"
	"github.com/spf13/cobra"
)

var createClipCmd = &cobra.Command{
	Use:   "clip",
	Short: "Create clips in an existing System Project",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf(
				"configuration is required before creating clips: %w",
				err,
			)
		}

		projects, err := filesystem.ListProjects(
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

		selectedProject, count, err := runCreateClipForm(
			projects,
		)
		if err != nil {
			return err
		}

		projectPath := filepath.Join(
			cfg.SystemProjectsRoot,
			selectedProject,
		)

		nextIndex, err := filesystem.NextClipIndex(
			projectPath,
		)
		if err != nil {
			return err
		}

		plan, err := project.BuildClipPlan(
			nextIndex,
			count,
		)
		if err != nil {
			return err
		}

		confirmed, err := confirmClipPlan(
			selectedProject,
			projectPath,
			plan,
		)
		if err != nil {
			return err
		}

		if !confirmed {
			fmt.Println("Clip creation cancelled.")
			return nil
		}

		executor := filesystem.NewExecutor()

		if err := executor.ExecuteDirectories(
			projectPath,
			plan.Directories,
		); err != nil {
			return err
		}

		fmt.Printf(
			"\nCreated %d clip(s) in %s.\n",
			count,
			selectedProject,
		)

		return nil
	},
}

func runCreateClipForm(
	projects []string,
) (string, int, error) {
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
				Validate(func(value string) error {
					count, err := strconv.Atoi(
						strings.TrimSpace(value),
					)
					if err != nil {
						return fmt.Errorf(
							"must be a number",
						)
					}

					if count < 1 {
						return fmt.Errorf(
							"must be greater than zero",
						)
					}

					return nil
				}),
		),
	)

	if err := form.Run(); err != nil {
		return "", 0, err
	}

	count, err := strconv.Atoi(
		strings.TrimSpace(countInput),
	)
	if err != nil {
		return "", 0, err
	}

	return selectedProject, count, nil
}

func confirmClipPlan(
	projectName string,
	projectPath string,
	plan project.ClipPlan,
) (bool, error) {
	var builder strings.Builder

	fmt.Fprintf(
		&builder,
		"Project: %s\n",
		projectName,
	)

	fmt.Fprintf(
		&builder,
		"Location: %s\n\n",
		projectPath,
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
				Title("Create these clips?").
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
	createCmd.AddCommand(createClipCmd)
}
