package dispatcher

import (
	"os"

	"github.com/brownhounds/migoro/utils"
)

func Make(n string) {
	if !utils.ValidateStringANU(n) {
		utils.Error("Migration name", "Only alphanumeric characters and underscores are allowed for migration name.")
		os.Exit(1)
	}

	utils.CreateMigrationFile(n)
}
