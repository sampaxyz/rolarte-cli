package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Rolarte CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Config")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
