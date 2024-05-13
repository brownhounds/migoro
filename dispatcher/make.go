package dispatcher

import (
	"migoro/utils"
	"os"
)

func Make(n string) {
	if !utils.ValidateStringANU(n) {
		utils.Error("Migration name", "Only alphanumeric characters and underscores are allowed for migration name.")
		os.Exit(1)
	}

	utils.CreateMigration(n)
}
