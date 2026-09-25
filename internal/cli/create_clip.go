package cli

import (
	"github.com/sampaxyz/rolarte-cli/internal/app"
	"github.com/sampaxyz/rolarte-cli/internal/interactive"
	"github.com/spf13/cobra"
)

var createClipCmd = &cobra.Command{
	Use:   "clip",
	Short: "Create clips in an existing System Project",

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {
		return interactive.CreateClip(
			app.New(),
		)
	},
}

func init() {
	createCmd.AddCommand(createClipCmd)
}
