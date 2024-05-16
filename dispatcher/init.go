package dispatcher

import (
	"migoro/adapters"
	"migoro/error_context"
	"migoro/types"
	"migoro/utils"
)

func initializeDatabase(adapter types.Adapter) {
	err, result := adapter.DatabaseExists()
	if err != nil {
		error_context.Context.SetError()
		return
	}

	if !result.Exists {
		adapter.CreateDatabase()
		utils.Success("Database Created", adapter.GetDatabaseName())
	} else {
		utils.Warning("Database already exists", adapter.GetDatabaseName())
	}
}

func initializeMigrationLog(adapter types.Adapter) {
	err, migrationTable := adapter.MigrationsLogExists()

	if err != nil {
		error_context.Context.SetError()
		return
	}

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
