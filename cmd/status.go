package cmd

import (
	"migoro/dispatcher"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get migration status",
	Long:  "\n" + `Print migration status to see which scripts have been applied. Files will be marked as APPLIED, NOT APPLIED, EMPTY FILE.`,
	Run: func(cmd *cobra.Command, args []string) {
		dispatcher.Status()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
