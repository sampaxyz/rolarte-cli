package cli

import (
	"fmt"

	"charm.land/huh/v2"
	"github.com/sampaxyz/rolarte-cli/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Rolarte CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		var systemProjectsRoot string

		currentConfig, err := config.Load()
		if err == nil {
			systemProjectsRoot = currentConfig.SystemProjectsRoot
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("System Projects Root").
					Description("Directory where Rolarte System Projects are stored.").
					Value(&systemProjectsRoot).
					Validate(func(value string) error {
						if value == "" {
							return fmt.Errorf("path cannot be empty")
						}

						return nil
					}),
			),
		)

		if err := form.Run(); err != nil {
			return err
		}

		cfg := config.Config{
			SystemProjectsRoot: systemProjectsRoot,
		}

		if err := config.Save(cfg); err != nil {
			return err
		}

		path, err := config.Path()
		if err != nil {
			return err
		}

		fmt.Printf("\nConfiguration saved to:\n%s\n", path)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
