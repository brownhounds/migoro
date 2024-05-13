package cmd

import (
	"migoro/dispatcher"

	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make [MIGRATION NAME]",
	Args:  cobra.ExactArgs(1),
	Short: "Make a new migration file",
	Long:  "\n" + `By default this command will create a file with comments where migration scripts can be placed, respectively for down and up method. Comments must be in original format.`,
	Run: func(cmd *cobra.Command, args []string) {
		dispatcher.Make(args[0])
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}
