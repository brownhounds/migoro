package dispatcher

import (
	"migoro/adapters"
	"migoro/types"
	"migoro/utils"
)

func initializeMigrationLog(adapter types.Adapter) {
	migrationTable := adapter.MigrationsLogExists()

	if !migrationTable.Exists {
		adapter.CreateMigrationsLog()
		utils.Success("Table Created", utils.Env("MIGRATION_TABLE")) // TODO: WHAAAAT??
	} else {
		utils.Warning("Table already exists", utils.Env("MIGRATION_TABLE"))
	}
}

func initializeDatabase(adapter types.Adapter) {
	if !adapter.DatabaseExists().Exists {
		adapter.CreateDatabase()
		utils.Success("Database Created", utils.Env("SQL_DB"))
	} else {
		utils.Warning("Database already exists", utils.Env("SQL_DB"))
	}
}

func Init() {
	adapter := adapters.Init()
	initializeDatabase(adapter)
	initializeMigrationLog(adapter)
}
