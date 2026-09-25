package cli

import (
	"github.com/sampaxyz/rolarte-cli/internal/app"
	"github.com/sampaxyz/rolarte-cli/internal/interactive"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Rolarte CLI",

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {
		return interactive.Configure(
			app.New(),
		)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
