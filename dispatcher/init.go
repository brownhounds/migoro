package dispatcher

import (
	"migoro/adapters"
	"migoro/types"
	"migoro/utils"
)

func initializeDatabase(adapter types.Adapter) {
	if !adapter.DatabaseExists().Exists {
		adapter.CreateDatabase()
		utils.Success("Database Created", adapter.GetDatabaseName())
	} else {
		utils.Warning("Database already exists", adapter.GetDatabaseName())
	}
}

func initializeMigrationLog(adapter types.Adapter) {
	migrationTable := adapter.MigrationsLogExists()

	if !migrationTable.Exists {
		adapter.CreateMigrationsLog()
		utils.Success("Table Created", adapter.GetMigrationTableName())
	} else {
		utils.Warning("Table already exists", adapter.GetMigrationTableName())
	}
}

func Init() {
	adapter := adapters.Init()
	initializeDatabase(adapter)
	initializeMigrationLog(adapter)
}
