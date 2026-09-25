package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createClipCmd = &cobra.Command{
	Use:   "clip",
	Short: "Create clips in an existing System Project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create clip")
	},
}

func init() {
	createCmd.AddCommand(createClipCmd)
}
