package dispatcher

import (
	"migoro/adapters"
	"migoro/query"
	"migoro/utils"
)

func Init() {
	adapter := adapters.Init()

	migrationTable := query.Exists(adapter.TableLogExistsQuery())

	if !migrationTable.Exists {
		query.Query(adapter.CreateLogTableQuery())
		utils.Success("Table Created", utils.Env(adapters.MIGRATION_TABLE))
	} else {
		utils.Warning("Table already exists", utils.Env(adapters.MIGRATION_TABLE))
	}
}
