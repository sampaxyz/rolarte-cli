package cli

import (
	"fmt"

	"github.com/sampaxyz/rolarte-cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Rolarte CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"Rolarte CLI %s\ncommit: %s\nbuilt: %s\n",
			version.Version,
			version.Commit,
			version.Date,
		)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
