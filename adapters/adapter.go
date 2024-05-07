package adapters

import (
	"fmt"
	"migoro/adapters/postgres"
	"migoro/adapters/sqlite"
	"migoro/types"
	"migoro/utils"
	"os"
)

const (
	SQL_DRIVER = "SQL_DRIVER"
	POSTGRES = "postgres"
	SQLITE3 = "sqlite3"
)


func Init() types.Adapter {
	adapter, err := resolveAdapter();
	if err != nil {
		utils.Error("Resolving Adapter", err.Error())
		os.Exit(1)
	}

	adapter.ValidateEnvironment()

	return adapter
}

func resolveAdapter() (types.Adapter, error) {
	switch utils.Env(SQL_DRIVER) {
	case POSTGRES:
		return postgres.Postgres{}, nil
	case SQLITE3:
		return sqlite.Sqlite{}, nil
	default:
		return nil, fmt.Errorf("could not resolve the adapter: %s", utils.Env(SQL_DRIVER))
	}
}
