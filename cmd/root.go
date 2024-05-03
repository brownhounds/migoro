package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/logrusorgru/aurora/v4"
)

var rootCmd = &cobra.Command{
	Use:   "migoro",
	Short: "A brief description of your application",
	Long:  fmt.Sprintf(`%s - Database migration manager%s%s`, aurora.Cyan("\nMigoro").String(), "\nv1.0", "\n\nAvailable Drivers:\n- Postgres"),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
