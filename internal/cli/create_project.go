package cli

import (
	"fmt"

	"github.com/sampaxyz/rolarte-cli/internal/project"
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create a new System Project",
	RunE: func(cmd *cobra.Command, args []string) error {
		request := project.CreateRequest{
			Name: "Outta Bubblegum",
			Subprojects: []project.Subproject{
				project.SubprojectMicrofiction,
				project.SubprojectSystemOverview,
				project.SubprojectPodcastGame,
				project.SubprojectPodcastRolafter,
				project.SubprojectClips,
			},
			InitialClips: 1,
		}

		plan, err := project.BuildPlan(request)
		if err != nil {
			return err
		}

		fmt.Printf("Project: %s\n\n", plan.Root)

		for _, directory := range plan.Directories {
			fmt.Printf("  %s/\n", directory)
		}

		return nil
	},
}

func init() {
	createCmd.AddCommand(createProjectCmd)
}
