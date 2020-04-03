package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Check version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
