package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "rolarte",
	Short: "Rolarte audiovisual production toolkit",
	Long:  `Rolarte CLI standardizes and automates audiovisual production workflows used by Rolarte.`,
}

func Execute() error {
	return rootCmd.Execute()
}
