package adapters

import (
	"fmt"
	"migoro/adapters/postgres"
	"migoro/adapters/sqlite"
	"migoro/types"
	"migoro/utils"
)

const (
	SQL_DRIVER = "SQL_DRIVER"
	POSTGRES   = "postgres"
	SQLITE3    = "sqlite3"
)

func UnsetAdapterErrorMessage(driver string) string {
	return fmt.Sprintf("Adapter is not set: consider setting environment variable - %s", driver)
}

func UnsupportedAdapterErrorMessage(driver string) string {
	return fmt.Sprintf("Unsupported adapter: %s", driver)
}

func Init() types.Adapter {
	adapter := resolveAdapter()
	adapter.ValidateEnvironment()
	return adapter
}

func resolveAdapter() types.Adapter {
	switch utils.Env(SQL_DRIVER) {
	case POSTGRES:
		return &postgres.Postgres{}
	case SQLITE3:
		return &sqlite.Sqlite{}
	default:
		if utils.Env(SQL_DRIVER) == "" {
			panic(UnsetAdapterErrorMessage(SQL_DRIVER))
		}

		panic(UnsupportedAdapterErrorMessage(utils.Env(SQL_DRIVER)))
	}
}
