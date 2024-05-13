package cmd

import (
	"migoro/dispatcher"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply migrations",
	Long:  "\n" + `Will apply all the migrations in the ascending order, marking all the migrated scripts with the same hash. In case of the rollback manager will attempt to run relevant rollback script for all the migrations applied during the execution of this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		dispatcher.Init()
		dispatcher.Migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
