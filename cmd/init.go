package cmd

import (
	"github.com/brownhounds/migoro/dispatcher"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize database table for migration log",
	Long:  "\n" + `Initialize database table for migration log. In case of Postgres this command will create the schema along with the table. Name of the table and schema can be configured in .env file.`,
	Run: func(cmd *cobra.Command, args []string) {
		dispatcher.Init()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
