package cli

import (
	"github.com/sampaxyz/rolarte-cli/internal/app"
	"github.com/sampaxyz/rolarte-cli/internal/interactive"
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create a new System Project",

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {
		return interactive.CreateProject(
			app.New(),
		)
	},
}

func init() {
	createCmd.AddCommand(createProjectCmd)
}
