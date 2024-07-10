package cmd

import (
	"fmt"
	"os"

	"github.com/brownhounds/migoro/version"

	"github.com/spf13/cobra"

	"github.com/logrusorgru/aurora/v4"
)

var rootCmd = &cobra.Command{
	Use:     "migoro",
	Short:   "CLI Database Migrator",
	Version: version.Ver,
	Long:    fmt.Sprintf(`%s - Database migration manager%s`, aurora.Cyan("\nMigoro").String(), "\n\nAvailable Drivers:\n- Postgres\n- SQLite3"),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
