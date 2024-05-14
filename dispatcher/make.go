package dispatcher

import (
	"migoro/error_context"
	"migoro/utils"
)

func Make(n string) {
	if !utils.ValidateStringANU(n) {
		error_context.Context.SetError()
		utils.Error("Migration name", "Only alphanumeric characters and underscores are allowed for migration name.")
		return
	}

	utils.CreateMigrationFile(n)
}
