package cli

import (
	"github.com/sampaxyz/rolarte-cli/internal/app"
	"github.com/sampaxyz/rolarte-cli/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rolarte",
	Short: "Rolarte audiovisual production toolkit",

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {
		return tui.Run(
			app.New(),
		)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
