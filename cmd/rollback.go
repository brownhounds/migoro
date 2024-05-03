package cmd

import (
	"migoro/dispatcher"

	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback latest migrations",
	Long:  "\n" + `Revert all the changes applied by the late 'migrate' command. The manager will attempt to run all the corresponding DOWN scripts. In case of missing or empty files rollback will be prevented.`,
	Run: func(cmd *cobra.Command, args []string) {
		dispatcher.Rollback()
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
}
